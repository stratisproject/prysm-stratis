package finality

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/altair/finality"
)

func TestMinimal_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
