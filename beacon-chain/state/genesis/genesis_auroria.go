//go:build !noAuroriaGenesis
// +build !noAuroriaGenesis

package genesis

import (
	_ "embed"

	"github.com/prysmaticlabs/prysm/v4/config/params"
)

var (
	//go:embed auroria.ssz.snappy
	auroriaRawSSZCompressed []byte // 2.7Mb
)

func init() {
	embeddedStates[params.AuroriaName] = &auroriaRawSSZCompressed
}
