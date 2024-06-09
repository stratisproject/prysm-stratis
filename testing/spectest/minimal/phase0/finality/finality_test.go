package finality

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/phase0/finality"
)

func TestMinimal_Phase0_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
