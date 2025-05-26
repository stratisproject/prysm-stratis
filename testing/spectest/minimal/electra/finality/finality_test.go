package finality

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/finality"
)

func TestMinimal_Electra_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
