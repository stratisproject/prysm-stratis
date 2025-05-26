package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/runtime/version"
	common "github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/operations"
)

func RunBlockHeaderTest(t *testing.T, config string) {
	common.RunBlockHeaderTest(t, config, version.String(version.Capella), sszToBlock, sszToState)
}
