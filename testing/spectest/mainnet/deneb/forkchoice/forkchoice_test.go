package forkchoice

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Deneb_Forkchoice(t *testing.T) {
	t.Skip("This will fail until we re-integrate proof verification")
	forkchoice.Run(t, "mainnet", version.Deneb)
}
