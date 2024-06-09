package ssz_static

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/ssz_static"
)

func TestMinimal_Deneb_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "minimal")
}
