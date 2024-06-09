package random

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/sanity"
)

func TestMinimal_Deneb_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "random/random/pyspec_tests")
}
