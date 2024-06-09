package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/bellatrix/operations"
)

func TestMainnet_Bellatrix_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "mainnet")
}
