package fork_helper

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/fork"
)

func TestMainnet_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "mainnet")
}
