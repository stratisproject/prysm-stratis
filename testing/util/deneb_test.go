package util

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestInclusionProofs(t *testing.T) {
	_, blobs := GenerateTestDenebBlockWithSidecar(t, [32]byte{}, 0, params.BeaconConfig().MaxBlobsPerBlock(0))
	for i := range blobs {
		require.NoError(t, blocks.VerifyKZGInclusionProof(blobs[i]))
	}
}
