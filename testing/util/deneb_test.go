package util

import (
	"testing"

	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestInclusionProofs(t *testing.T) {
	_, blobs := GenerateTestDenebBlockWithSidecar(t, [32]byte{}, 0, fieldparams.MaxBlobsPerBlock)
	for i := range blobs {
		require.NoError(t, blocks.VerifyKZGInclusionProof(blobs[i]))
	}
}
