package keeper

import (
	"github.com/hengmengsroin/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
