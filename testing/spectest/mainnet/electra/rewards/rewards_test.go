package rewards

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/rewards"
)

func TestMainnet_Electra_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "mainnet")
}
