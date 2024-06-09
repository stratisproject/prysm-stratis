package kv

import (
	"context"
	"testing"

	testpb "github.com/stratisproject/prysm-stratis/proto/testing"
	"github.com/stratisproject/prysm-stratis/testing/require"
)

func Test_encode_handlesNilFromFunction(t *testing.T) {
	foo := func() *testpb.Puzzle {
		return nil
	}
	_, err := encode(context.Background(), foo())
	require.ErrorContains(t, "cannot encode nil message", err)
}
