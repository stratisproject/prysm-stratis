package ssz_static

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/capella/ssz_static"
)

func TestMinimal_Capella_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "minimal")
}
