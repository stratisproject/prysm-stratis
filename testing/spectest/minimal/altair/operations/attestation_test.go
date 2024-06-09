package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
