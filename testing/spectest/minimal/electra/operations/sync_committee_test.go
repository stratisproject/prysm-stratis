package operations

import (
	"testing"

	"github.com/stratisproject/prysm-stratis/testing/spectest/shared/electra/operations"
)

func TestMinimal_Electra_Operations_SyncCommittee(t *testing.T) {
	operations.RunSyncCommitteeTest(t, "minimal")
}
