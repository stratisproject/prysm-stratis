package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
