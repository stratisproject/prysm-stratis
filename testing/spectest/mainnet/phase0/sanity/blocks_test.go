package sanity

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/phase0/sanity"
)

func TestMainnet_Phase0_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "sanity/blocks/pyspec_tests")
}
