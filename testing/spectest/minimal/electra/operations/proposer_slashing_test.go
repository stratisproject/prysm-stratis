package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/operations"
)

func TestMinimal_Electra_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}
