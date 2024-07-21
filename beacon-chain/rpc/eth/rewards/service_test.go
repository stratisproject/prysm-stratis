package rewards

import (
	"context"
	"testing"

	"github.com/stratisproject/prysm-stratis/beacon-chain/core/transition"
	dbutil "github.com/stratisproject/prysm-stratis/beacon-chain/db/testing"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/testing/assert"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func TestGetStateForRewards_NextSlotCacheHit(t *testing.T) {
	ctx := context.Background()
	db := dbutil.SetupDB(t)

	st, err := util.NewBeaconStateDeneb()
	require.NoError(t, err)
	b := util.HydrateSignedBeaconBlockDeneb(util.NewBeaconBlockDeneb())
	parent, err := blocks.NewSignedBeaconBlock(b)
	require.NoError(t, err)
	require.NoError(t, db.SaveBlock(ctx, parent))

	r, err := parent.Block().HashTreeRoot()
	require.NoError(t, err)
	require.NoError(t, transition.UpdateNextSlotCache(ctx, r[:], st))

	s := &BlockRewardService{
		Replayer: nil, // setting to nil because replayer must not be invoked
		DB:       db,
	}
	b = util.HydrateSignedBeaconBlockDeneb(util.NewBeaconBlockDeneb())
	sbb, err := blocks.NewSignedBeaconBlock(b)
	require.NoError(t, err)
	sbb.SetSlot(parent.Block().Slot() + 1)
	result, err := s.GetStateForRewards(ctx, sbb.Block())
	require.NoError(t, err)
	_, lcs := transition.LastCachedState()
	expected, err := lcs.HashTreeRoot(ctx)
	require.NoError(t, err)
	actual, err := result.HashTreeRoot(ctx)
	require.NoError(t, err)
	assert.DeepEqual(t, expected, actual)
}
