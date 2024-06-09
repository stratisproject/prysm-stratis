package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
