package fork

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/fork"
)

func TestMinimal_Capella_UpgradeToCapella(t *testing.T) {
	fork.RunUpgradeToCapella(t, "minimal")
}
