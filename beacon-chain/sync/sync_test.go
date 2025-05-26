package sync

import (
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/cmd/beacon-chain/flags"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)

	resetFlags := flags.Get()
	flags.Init(&flags.GlobalFlags{
		BlockBatchLimit:            64,
		BlockBatchLimitBurstFactor: 10,
		BlobBatchLimit:             32,
		BlobBatchLimitBurstFactor:  2,
	})
	defer func() {
		flags.Init(resetFlags)
	}()
	os.Exit(m.Run())
}
