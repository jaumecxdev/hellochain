package keeper

import (
	"hellochain/x/hellochain/types"
)

var _ types.QueryServer = Keeper{}
