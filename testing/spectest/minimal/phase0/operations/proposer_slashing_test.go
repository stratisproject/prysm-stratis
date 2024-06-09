package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/phase0/operations"
)

func TestMinimal_Phase0_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}
