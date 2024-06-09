package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_AttesterSlashing(t *testing.T) {
	operations.RunAttesterSlashingTest(t, "mainnet")
}
