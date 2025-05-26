package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_Consolidation(t *testing.T) {
	operations.RunConsolidationTest(t, "mainnet")
}
