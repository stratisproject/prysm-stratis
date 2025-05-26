package fork_helper

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/fork"
)

func TestMainnet_UpgradeToElectra(t *testing.T) {
	fork.RunUpgradeToElectra(t, "mainnet")
}
