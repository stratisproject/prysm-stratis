package builder

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

type Server struct {
	FinalizationFetcher   blockchain.FinalizationFetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	Stater                lookup.Stater
}
