package epoch_processing

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/epoch_processing"
)

func TestMinimal_electra_EpochProcessing_EffectiveBalanceUpdates(t *testing.T) {
	epoch_processing.RunEffectiveBalanceUpdatesTests(t, "minimal")
}
