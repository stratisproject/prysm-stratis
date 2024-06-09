package validator

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/core"
)

type Server struct {
	CoreService *core.Service
}
