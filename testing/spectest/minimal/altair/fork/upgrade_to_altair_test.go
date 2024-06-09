package fork

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/altair/fork"
)

func TestMinimal_Altair_UpgradeToAltair(t *testing.T) {
	fork.RunUpgradeToAltair(t, "minimal")
}
