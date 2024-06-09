//go:build !noAuroriaGenesis
// +build !noAuroriaGenesis

package genesis

import (
	_ "embed"

	"github.com/stratisproject/prysm-stratis/config/params"
)

var (
	//go:embed auroria.ssz.snappy
	auroriaRawSSZCompressed []byte // 2.7Mb
)

func init() {
	embeddedStates[params.AuroriaName] = &auroriaRawSSZCompressed
}
