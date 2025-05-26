package sanity

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/sanity"
)

func TestMainnet_Electra_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "sanity/blocks/pyspec_tests")
}
