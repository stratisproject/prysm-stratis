package endtoend

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/testing/endtoend/types"
)

func TestEndToEnd_MinimalConfig(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Phase0, version.Deneb, params.E2ETestConfig()), types.WithCheckpointSync())
	r.run()
}
