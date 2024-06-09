package random

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/sanity"
)

func TestMinimal_Capella_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "random/random/pyspec_tests")
}
