package finality

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/finality"
)

func TestMainnet_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
