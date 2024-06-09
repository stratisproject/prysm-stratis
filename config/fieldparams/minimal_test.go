//go:build minimal

package field_params_test

import (
	"testing"

	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	"github.com/stratisproject/prysm-stratis/config/params"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestFieldParametersValues(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	min := params.MinimalSpecConfig().Copy()
	params.OverrideBeaconConfig(min)
	require.Equal(t, "minimal", fieldparams.Preset)
	testFieldParametersMatchConfig(t)
}
