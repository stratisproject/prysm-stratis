package rewards

import (
	"context"
	"net/http"
	"strconv"

	"github.com/stratisproject/prysm-stratis/api/server/structs"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/altair"
	coreblocks "github.com/stratisproject/prysm-stratis/beacon-chain/core/blocks"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/transition"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/validators"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen"
	consensusblocks "github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	"github.com/stratisproject/prysm-stratis/network/httputil"
	"github.com/stratisproject/prysm-stratis/time/slots"
)

// BlockRewardsFetcher is a interface that provides access to reward related responses
type BlockRewardsFetcher interface {
	GetBlockRewardsData(context.Context, interfaces.ReadOnlyBeaconBlock) (*structs.BlockRewards, *httputil.DefaultJsonError)
	GetStateForRewards(context.Context, interfaces.ReadOnlyBeaconBlock) (state.BeaconState, *httputil.DefaultJsonError)
}

// BlockRewardService implements BlockRewardsFetcher and can be declared to access the underlying functions
type BlockRewardService struct {
	Replayer stategen.ReplayerBuilder
	DB       db.HeadAccessDatabase
}

// GetBlockRewardsData returns the BlockRewards object which is used for the BlockRewardsResponse and ProduceBlockV3.
// Rewards are denominated in Gwei.
func (rs *BlockRewardService) GetBlockRewardsData(ctx context.Context, blk interfaces.ReadOnlyBeaconBlock) (*structs.BlockRewards, *httputil.DefaultJsonError) {
	if blk == nil || blk.IsNil() {
		return nil, &httputil.DefaultJsonError{
			Message: consensusblocks.ErrNilBeaconBlock.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	st, httpErr := rs.GetStateForRewards(ctx, blk)
	if httpErr != nil {
		return nil, httpErr
	}

	proposerIndex := blk.ProposerIndex()
	initBalance, err := st.BalanceAtIndex(proposerIndex)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get proposer's balance: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	st, err = altair.ProcessAttestationsNoVerifySignature(ctx, st, blk)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get attestation rewards: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	attBalance, err := st.BalanceAtIndex(proposerIndex)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get proposer's balance: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	st, err = coreblocks.ProcessAttesterSlashings(ctx, st, blk.Body().AttesterSlashings(), validators.SlashValidator)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get attester slashing rewards: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	attSlashingsBalance, err := st.BalanceAtIndex(proposerIndex)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get proposer's balance: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	st, err = coreblocks.ProcessProposerSlashings(ctx, st, blk.Body().ProposerSlashings(), validators.SlashValidator)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get proposer slashing rewards: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	proposerSlashingsBalance, err := st.BalanceAtIndex(proposerIndex)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get proposer's balance: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	sa, err := blk.Body().SyncAggregate()
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get sync aggregate: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	var syncCommitteeReward uint64
	_, syncCommitteeReward, err = altair.ProcessSyncAggregate(ctx, st, sa)
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get sync aggregate rewards: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return &structs.BlockRewards{
		ProposerIndex:     strconv.FormatUint(uint64(proposerIndex), 10),
		Total:             strconv.FormatUint(proposerSlashingsBalance-initBalance+syncCommitteeReward, 10),
		Attestations:      strconv.FormatUint(attBalance-initBalance, 10),
		SyncAggregate:     strconv.FormatUint(syncCommitteeReward, 10),
		ProposerSlashings: strconv.FormatUint(proposerSlashingsBalance-attSlashingsBalance, 10),
		AttesterSlashings: strconv.FormatUint(attSlashingsBalance-attBalance, 10),
	}, nil
}

// GetStateForRewards returns the state replayed up to the block's slot
func (rs *BlockRewardService) GetStateForRewards(ctx context.Context, blk interfaces.ReadOnlyBeaconBlock) (state.BeaconState, *httputil.DefaultJsonError) {
	// We want to run several block processing functions that update the proposer's balance.
	// This will allow us to calculate proposer rewards for each operation (atts, slashings etc).
	// To do this, we replay the state up to the block's slot, but before processing the block.

	// Try getting the state from the next slot cache first.
	_, prevSlotRoots, err := rs.DB.BlockRootsBySlot(ctx, slots.PrevSlot(blk.Slot()))
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get roots for previous slot: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	for _, r := range prevSlotRoots {
		s := transition.NextSlotState(r[:], blk.Slot())
		if s != nil {
			return s, nil
		}
	}

	st, err := rs.Replayer.ReplayerForSlot(slots.PrevSlot(blk.Slot())).ReplayToSlot(ctx, blk.Slot())
	if err != nil {
		return nil, &httputil.DefaultJsonError{
			Message: "Could not get state: " + err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return st, nil
}
