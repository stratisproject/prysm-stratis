package scorers_test

import (
	"io"
	"math"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/beacon-chain/p2p/peers/scorers"
	"github.com/stratisproject/prysm-stratis/cmd/beacon-chain/flags"
	"github.com/stratisproject/prysm-stratis/config/features"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)

	resetCfg := features.InitWithReset(&features.Flags{
		EnablePeerScorer: true,
	})
	defer resetCfg()

	resetFlags := flags.Get()
	flags.Init(&flags.GlobalFlags{
		BlockBatchLimit:            64,
		BlockBatchLimitBurstFactor: 10,
	})
	defer func() {
		flags.Init(resetFlags)
	}()
	m.Run()
}

// roundScore returns score rounded in accordance with the score manager's rounding factor.
func roundScore(score float64) float64 {
	return math.Round(score*scorers.ScoreRoundingFactor) / scorers.ScoreRoundingFactor
}
