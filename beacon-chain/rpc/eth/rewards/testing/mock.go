package testing

import (
	"context"

	"github.com/stratisproject/prysm-stratis/api/server/structs"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	"github.com/stratisproject/prysm-stratis/network/httputil"
)

type MockBlockRewardFetcher struct {
	Rewards *structs.BlockRewards
	Error   *httputil.DefaultJsonError
	State   state.BeaconState
}

func (m *MockBlockRewardFetcher) GetBlockRewardsData(_ context.Context, _ interfaces.ReadOnlyBeaconBlock) (*structs.BlockRewards, *httputil.DefaultJsonError) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Rewards, nil
}

func (m *MockBlockRewardFetcher) GetStateForRewards(_ context.Context, _ interfaces.ReadOnlyBeaconBlock) (state.BeaconState, *httputil.DefaultJsonError) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.State, nil
}
