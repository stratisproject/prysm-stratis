package beacon

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	beacondb "github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen"
	"github.com/stratisproject/prysm-stratis/beacon-chain/sync"
)

type Server struct {
	SyncChecker           sync.Checker
	HeadFetcher           blockchain.HeadFetcher
	TimeFetcher           blockchain.TimeFetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	CanonicalHistory      *stategen.CanonicalHistory
	BeaconDB              beacondb.ReadOnlyDatabase
	Stater                lookup.Stater
	ChainInfoFetcher      blockchain.ChainInfoFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
}
