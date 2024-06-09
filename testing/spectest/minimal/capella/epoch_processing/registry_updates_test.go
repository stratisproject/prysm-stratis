package epoch_processing

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/epoch_processing"
)

func TestMinimal_Capella_EpochProcessing_ResetRegistryUpdates(t *testing.T) {
	epoch_processing.RunRegistryUpdatesTests(t, "minimal")
}
