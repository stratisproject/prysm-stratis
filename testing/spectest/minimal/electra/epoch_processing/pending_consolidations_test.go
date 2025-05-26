package epoch_processing

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/epoch_processing"
)

func TestMinimal_Electra_EpochProcessing_PendingConsolidations(t *testing.T) {
	epoch_processing.RunPendingConsolidationsTests(t, "minimal")
}
