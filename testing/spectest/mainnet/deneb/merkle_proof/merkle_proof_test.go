package merkle_proof

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/merkle_proof"
)

func TestMainnet_Deneb_MerkleProof(t *testing.T) {
	merkle_proof.RunMerkleProofTests(t, "mainnet")
}
