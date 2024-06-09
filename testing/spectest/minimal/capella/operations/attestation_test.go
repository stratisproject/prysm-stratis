package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
