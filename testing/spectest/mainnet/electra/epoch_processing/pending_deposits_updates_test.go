package epoch_processing

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/epoch_processing"
)

func TestMainnet_Electra_EpochProcessing_PendingDeposits(t *testing.T) {
	epoch_processing.RunPendingDepositsTests(t, "mainnet")
}
