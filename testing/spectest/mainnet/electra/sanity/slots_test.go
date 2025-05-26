package sanity

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/sanity"
)

func TestMainnet_Electra_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "mainnet")
}
