package db

import "github.com/stratisproject/prysm-stratis/beacon-chain/db/kv"

var _ Database = (*kv.Store)(nil)
