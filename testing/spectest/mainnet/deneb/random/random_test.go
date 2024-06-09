package random

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/sanity"
)

func TestMainnet_Deneb_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}
