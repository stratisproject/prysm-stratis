package epoch_processing

import (
	"path"
	"testing"

	"github.com/stratisproject/prysm-stratis/beacon-chain/core/electra"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/helpers"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	"github.com/stratisproject/prysm-stratis/testing/require"
	"github.com/stratisproject/prysm-stratis/testing/spectest/utils"
)

// RunSlashingsTests executes "epoch_processing/slashings" tests.
func RunSlashingsTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "electra", "epoch_processing/slashings/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			helpers.ClearCache()
			RunEpochOperationTest(t, folderPath, processSlashingsWrapper)
		})
	}
}

func processSlashingsWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	require.NoError(t, electra.ProcessSlashings(st), "Could not process slashings")
	return st, nil
}
