package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/runtime/version"
	common "github.com/stratisproject/prysm-stratis/testing/spectest/shared/common/operations"
)

func RunExecutionPayloadTest(t *testing.T, config string) {
	common.RunExecutionPayloadTest(t, config, version.String(version.Capella), sszToBlockBody, sszToState)
}
