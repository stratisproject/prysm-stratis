package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/operations"
)

func TestMinimal_Electra_Operations_AttesterSlashing(t *testing.T) {
	operations.RunAttesterSlashingTest(t, "minimal")
}
