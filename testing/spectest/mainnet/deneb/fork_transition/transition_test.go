package fork_transition

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/fork"
)

func TestMainnet_Deneb_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
