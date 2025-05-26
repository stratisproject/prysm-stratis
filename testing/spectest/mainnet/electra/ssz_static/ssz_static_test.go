package ssz_static

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/ssz_static"
)

func TestMainnet_Electra_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}
