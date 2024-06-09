package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
