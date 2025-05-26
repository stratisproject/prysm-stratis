package merkle_proof

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/merkle_proof"
)

func TestMainnet_Electra_MerkleProof(t *testing.T) {
	merkle_proof.RunMerkleProofTests(t, "mainnet")
}
