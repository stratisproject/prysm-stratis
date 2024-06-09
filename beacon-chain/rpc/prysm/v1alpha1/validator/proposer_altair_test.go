package validator

import (
	"context"
	"testing"

	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/util"
)

func TestServer_SetSyncAggregate_EmptyCase(t *testing.T) {
	b, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlockAltair())
	require.NoError(t, err)
	s := &Server{} // Sever is not initialized with sync committee pool.
	s.setSyncAggregate(context.Background(), b)
	agg, err := b.Block().Body().SyncAggregate()
	require.NoError(t, err)

	emptySig := [96]byte{0xC0}
	want := &ethpb.SyncAggregate{
		SyncCommitteeBits:      make([]byte, params.BeaconConfig().SyncCommitteeSize/8),
		SyncCommitteeSignature: emptySig[:],
	}
	require.DeepEqual(t, want, agg)
}
