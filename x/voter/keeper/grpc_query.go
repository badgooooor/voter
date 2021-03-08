package keeper

import (
	"github.com/borbier/voter/x/voter/types"
)

var _ types.QueryServer = Keeper{}
