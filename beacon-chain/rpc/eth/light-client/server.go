package lightclient

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	lightClient "github.com/stratisproject/prysm-stratis/beacon-chain/core/light-client"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker          lookup.Blocker
	Stater           lookup.Stater
	HeadFetcher      blockchain.HeadFetcher
	ChainInfoFetcher blockchain.ChainInfoFetcher
	BeaconDB         db.HeadAccessDatabase
	LCStore          *lightClient.Store
}
