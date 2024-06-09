package lightclient

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker     lookup.Blocker
	Stater      lookup.Stater
	HeadFetcher blockchain.HeadFetcher
}
