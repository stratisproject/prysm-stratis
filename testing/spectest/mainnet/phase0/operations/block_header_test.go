package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/phase0/operations"
)

func TestMainnet_Phase0_Operations_BlockHeader(t *testing.T) {
	operations.RunBlockHeaderTest(t, "mainnet")
}
