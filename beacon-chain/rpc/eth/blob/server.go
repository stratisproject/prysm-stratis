package blob

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker lookup.Blocker
}
