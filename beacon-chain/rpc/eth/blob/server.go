package blob

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker               lookup.Blocker
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
	TimeFetcher           blockchain.TimeFetcher
}
