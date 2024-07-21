//go:build fuzz

package helpers

import "github.com/stratisproject/prysm-stratis/beacon-chain/cache"

func CommitteeCache() *cache.FakeCommitteeCache {
	return committeeCache
}

func SyncCommitteeCache() *cache.FakeSyncCommitteeCache {
	return syncCommitteeCache
}

func ProposerIndicesCache() *cache.FakeProposerIndicesCache {
	return proposerIndicesCache
}
