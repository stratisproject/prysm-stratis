package operations

import (
	"context"
	"testing"

	"github.com/stratisproject/prysm-stratis/beacon-chain/core/blocks"
	v "github.com/stratisproject/prysm-stratis/beacon-chain/core/validators"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
)

func RunProposerSlashingTest(t *testing.T, config string, fork string, block blockWithSSZObject, sszToState SSZToState) {
	runSlashingTest(t, config, fork, "proposer_slashing", block, sszToState, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
		return blocks.ProcessProposerSlashings(ctx, s, b.Block().Body().ProposerSlashings(), v.SlashValidator)
	})
}
