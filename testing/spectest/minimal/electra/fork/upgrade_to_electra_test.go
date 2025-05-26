package fork

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/fork"
)

func TestMinimal_UpgradeToElectra(t *testing.T) {
	fork.RunUpgradeToElectra(t, "minimal")
}
