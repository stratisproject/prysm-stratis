package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_BLSToExecutionChange(t *testing.T) {
	operations.RunBLSToExecutionChangeTest(t, "minimal")
}
