package fork

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/fulu/fork"
)

func TestMainnet_UpgradeToFulu(t *testing.T) {
	fork.RunUpgradeToFulu(t, "mainnet")
}
