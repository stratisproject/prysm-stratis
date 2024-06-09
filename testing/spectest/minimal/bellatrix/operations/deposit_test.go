package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
