package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "mainnet")
}
