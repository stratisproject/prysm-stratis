package fork

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/fork"
)

func TestMinimal_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "minimal")
}
