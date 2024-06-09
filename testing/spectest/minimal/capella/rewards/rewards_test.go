package rewards

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/rewards"
)

func TestMinimal_Capella_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "minimal")
}
