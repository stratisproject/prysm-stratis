package shared

import (
	"net/http"

	"github.com/stratisproject/prysm-stratis/beacon-chain/rpc/lookup"
	"github.com/stratisproject/prysm-stratis/consensus-types/blocks"
	"github.com/stratisproject/prysm-stratis/consensus-types/interfaces"
	"github.com/stratisproject/prysm-stratis/network/httputil"
)

// WriteStateFetchError writes an appropriate error based on the supplied argument.
// The argument error should be a result of fetching state.
func WriteStateFetchError(w http.ResponseWriter, err error) {
	if _, ok := err.(*lookup.StateNotFoundError); ok {
		httputil.HandleError(w, "State not found", http.StatusNotFound)
		return
	}
	if parseErr, ok := err.(*lookup.StateIdParseError); ok {
		httputil.HandleError(w, "Invalid state ID: "+parseErr.Error(), http.StatusBadRequest)
		return
	}
	httputil.HandleError(w, "Could not get state: "+err.Error(), http.StatusInternalServerError)
}

// WriteBlockFetchError writes an appropriate error based on the supplied argument.
// The argument error should be a result of fetching block.
func WriteBlockFetchError(w http.ResponseWriter, blk interfaces.ReadOnlySignedBeaconBlock, err error) bool {
	if invalidBlockIdErr, ok := err.(*lookup.BlockIdParseError); ok {
		httputil.HandleError(w, "Invalid block ID: "+invalidBlockIdErr.Error(), http.StatusBadRequest)
		return false
	}
	if err != nil {
		httputil.HandleError(w, "Could not get block from block ID: "+err.Error(), http.StatusInternalServerError)
		return false
	}
	if err = blocks.BeaconBlockIsNil(blk); err != nil {
		httputil.HandleError(w, "Could not find requested block: "+err.Error(), http.StatusNotFound)
		return false
	}
	return true
}
