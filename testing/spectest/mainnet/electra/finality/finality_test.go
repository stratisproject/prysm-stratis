package finality

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/finality"
)

func TestMainnet_Electra_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
