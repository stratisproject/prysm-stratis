package validator

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/core"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

type Server struct {
	BeaconDB            db.ReadOnlyDatabase
	Stater              lookup.Stater
	CanonicalFetcher    blockchain.CanonicalFetcher
	FinalizationFetcher blockchain.FinalizationFetcher
	ChainInfoFetcher    blockchain.ChainInfoFetcher
	CoreService         *core.Service
}
