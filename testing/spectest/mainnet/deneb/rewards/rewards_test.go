package rewards

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/rewards"
)

func TestMainnet_Deneb_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "mainnet")
}
