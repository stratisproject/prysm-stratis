package epoch_processing

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/epoch_processing"
)

func TestMainnet_Deneb_EpochProcessing_InactivityUpdates(t *testing.T) {
	epoch_processing.RunInactivityUpdatesTest(t, "mainnet")
}
