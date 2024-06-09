package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_SyncCommittee(t *testing.T) {
	operations.RunSyncCommitteeTest(t, "mainnet")
}
