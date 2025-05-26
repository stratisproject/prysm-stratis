package forkchoice

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/forkchoice"
)

func TestMinimal_Electra_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "minimal", version.Electra)
}
