package accounts

import (
	"encoding/json"
	"fmt"
	"path"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stratisproject/prysm-stratis/api/server/structs"
	"github.com/stratisproject/prysm-stratis/build/bazel"
	fieldparams "github.com/stratisproject/prysm-stratis/config/fieldparams"
	"github.com/stratisproject/prysm-stratis/encoding/bytesutil"
	"github.com/stratisproject/prysm-stratis/io/file"
	eth "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
	"github.com/stratisproject/prysm-stratis/testing/assert"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func TestDisplayExitInfo(t *testing.T) {
	logHook := test.NewGlobal()
	key := []byte("0x123456")
	displayExitInfo([][]byte{key}, []string{string(key)})
	assert.LogsContain(t, logHook, "https://beaconcha.in/validator/3078313233343536")
}

func TestDisplayExitInfo_NoKeys(t *testing.T) {
	logHook := test.NewGlobal()
	displayExitInfo([][]byte{}, []string{})
	assert.LogsContain(t, logHook, "No successful voluntary exits")
}

func TestPrepareAllKeys(t *testing.T) {
	key1 := bytesutil.ToBytes48([]byte("key1"))
	key2 := bytesutil.ToBytes48([]byte("key2"))
	raw, formatted := prepareAllKeys([][fieldparams.BLSPubkeyLength]byte{key1, key2})
	require.Equal(t, 2, len(raw))
	require.Equal(t, 2, len(formatted))
	assert.DeepEqual(t, bytesutil.ToBytes48([]byte{107, 101, 121, 49}), bytesutil.ToBytes48(raw[0]))
	assert.DeepEqual(t, bytesutil.ToBytes48([]byte{107, 101, 121, 50}), bytesutil.ToBytes48(raw[1]))
	assert.Equal(t, "0x6b6579310000", formatted[0])
	assert.Equal(t, "0x6b6579320000", formatted[1])
}

func TestWriteSignedVoluntaryExitJSON(t *testing.T) {
	sve := &eth.SignedVoluntaryExit{
		Exit: &eth.VoluntaryExit{
			Epoch:          5,
			ValidatorIndex: 300,
		},
		Signature: []byte{0x01, 0x02},
	}

	output := path.Join(bazel.TestTmpDir(), "TestWriteSignedVoluntaryExitJSON")
	require.NoError(t, writeSignedVoluntaryExitJSON(sve, output))

	b, err := file.ReadFileAsBytes(path.Join(output, "validator-exit-300.json"))
	require.NoError(t, err)

	svej := &structs.SignedVoluntaryExit{}
	require.NoError(t, json.Unmarshal(b, svej))

	require.Equal(t, fmt.Sprintf("%d", sve.Exit.Epoch), svej.Message.Epoch)
	require.Equal(t, fmt.Sprintf("%d", sve.Exit.ValidatorIndex), svej.Message.ValidatorIndex)
	require.Equal(t, "0x0102", svej.Signature)
}
