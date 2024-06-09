package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/operations"
)

func TestMinimal_Deneb_Operations_BLSToExecutionChange(t *testing.T) {
	operations.RunBLSToExecutionChangeTest(t, "minimal")
}
