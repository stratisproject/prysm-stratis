package core

import (
	"github.com/stratisproject/prysm-stratis/beacon-chain/blockchain"
	"github.com/stratisproject/prysm-stratis/beacon-chain/cache"
	opfeed "github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/operation"
	"github.com/stratisproject/prysm-stratis/beacon-chain/operations/synccommittee"
	"github.com/stratisproject/prysm-stratis/beacon-chain/p2p"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state/stategen"
	"github.com/stratisproject/prysm-stratis/beacon-chain/sync"
)

type Service struct {
	HeadFetcher           blockchain.HeadFetcher
	FinalizedFetcher      blockchain.FinalizationFetcher
	GenesisTimeFetcher    blockchain.TimeFetcher
	SyncChecker           sync.Checker
	Broadcaster           p2p.Broadcaster
	SyncCommitteePool     synccommittee.Pool
	OperationNotifier     opfeed.Notifier
	AttestationCache      *cache.AttestationCache
	StateGen              stategen.StateManager
	P2P                   p2p.Broadcaster
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
}
