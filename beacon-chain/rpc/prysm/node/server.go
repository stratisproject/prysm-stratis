package node

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/db"
	"github.com/stratisproject/prysm-stratis/beacon-chain/execution"
	"github.com/stratisproject/prysm-stratis/beacon-chain/p2p"
	"github.com/stratisproject/prysm-stratis/beacon-chain/sync"
)

type Server struct {
	SyncChecker               sync.Checker
	OptimisticModeFetcher     blockchain.OptimisticModeFetcher
	BeaconDB                  db.ReadOnlyDatabase
	PeersFetcher              p2p.PeersProvider
	PeerManager               p2p.PeerManager
	MetadataProvider          p2p.MetadataProvider
	GenesisTimeFetcher        blockchain.TimeFetcher
	HeadFetcher               blockchain.HeadFetcher
	ExecutionChainInfoFetcher execution.ChainInfoFetcher
}
