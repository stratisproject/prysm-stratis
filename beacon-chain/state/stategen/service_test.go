package stategen

import (
	"context"
	"testing"

	testDB "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	doublylinkedtree "github.com/stratisproject/prysm-stratis/beacon-chain/forkchoice/doubly-linked-tree"
	"github.com/stratisproject/prysm-stratis/config/params"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/assert"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func TestResume(t *testing.T) {
	ctx := context.Background()
	beaconDB := testDB.SetupDB(t)

	service := New(beaconDB, doublylinkedtree.New())
	b := util.NewBeaconBlock()
	util.SaveBlock(t, ctx, service.beaconDB, b)
	root, err := b.Block.HashTreeRoot()
	require.NoError(t, err)
	beaconState, _ := util.DeterministicGenesisState(t, 32)
	require.NoError(t, beaconState.SetSlot(params.BeaconConfig().SlotsPerEpoch))
	require.NoError(t, service.beaconDB.SaveState(ctx, beaconState, root))
	require.NoError(t, service.beaconDB.SaveGenesisBlockRoot(ctx, root))
	require.NoError(t, service.beaconDB.SaveFinalizedCheckpoint(ctx, &ethpb.Checkpoint{Root: root[:]}))

	resumeState, err := service.Resume(ctx, beaconState)
	require.NoError(t, err)
	require.DeepSSZEqual(t, beaconState.ToProtoUnsafe(), resumeState.ToProtoUnsafe())
	assert.Equal(t, params.BeaconConfig().SlotsPerEpoch, service.finalizedInfo.slot, "Did not get watned slot")
	assert.Equal(t, service.finalizedInfo.root, root, "Did not get wanted root")
	assert.NotNil(t, service.finalizedState(), "Wanted a non nil finalized state")
}
