package kv

import (
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/config/params"
)

func init() {
	// Override network name so that hardcoded genesis files are not loaded.
	if err := params.SetActive(params.MainnetTestConfig()); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)
	os.Exit(m.Run())
}
