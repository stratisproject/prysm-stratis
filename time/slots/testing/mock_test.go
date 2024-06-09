package testing

import (
	"github.com/stratisproject/prysm-stratis/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
