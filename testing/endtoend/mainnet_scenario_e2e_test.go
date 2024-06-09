package endtoend

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/runtime/version"
	"github.com/stratisproject/prysm-stratis/testing/endtoend/types"
)

func TestEndToEnd_MultiScenarioRun_Multiclient(t *testing.T) {
	runner := e2eMainnet(t, false, true, types.InitForkCfg(version.Phase0, version.Deneb, params.E2EMainnetTestConfig()), types.WithEpochs(24))
	runner.config.Evaluators = scenarioEvalsMulti()
	runner.config.EvalInterceptor = runner.multiScenarioMulticlient
	runner.scenarioRunner()
}
