package ssz_static

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/altair/ssz_static"
)

func TestMainnet_Altair_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}
