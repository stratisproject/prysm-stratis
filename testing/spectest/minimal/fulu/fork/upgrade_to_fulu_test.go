package fork

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/fulu/fork"
)

func TestMinimal_UpgradeToFulu(t *testing.T) {
	fork.RunUpgradeToFulu(t, "minimal")
}
