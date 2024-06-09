package forkchoice

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/forkchoice"
)

func TestMinimal_Deneb_Forkchoice(t *testing.T) {
	t.Skip("blocked by go-kzg-4844 minimal trusted setup")
	forkchoice.Run(t, "minimal", version.Deneb)
}
