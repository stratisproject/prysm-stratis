package fork_helper

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/fork"
)

func TestMainnet_Capella_UpgradeToCapella(t *testing.T) {
	fork.RunUpgradeToCapella(t, "mainnet")
}
