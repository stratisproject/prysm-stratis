package finality

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/finality"
)

func TestMinimal_Capella_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
