package fork_transition

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/bellatrix/fork"
)

func TestMinimal_Bellatrix_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
