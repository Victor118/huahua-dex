package keeper

import (
	"github.com/ChihuahuaChain/chihuahua/x/dex/types"
)

var _ types.QueryServer = Keeper{}
