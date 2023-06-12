// Copyright 2023 Origo Foundation
// This file is part of the Origo Network packages.
//
// Origo is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Origo packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Origo packages. If not, see https://github.com/OrigoTeam/origo/blob/main/LICENSE

package epochs

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/OrigoTeam/origo/x/epochs/keeper"
	"github.com/OrigoTeam/origo/x/epochs/types"
)

// InitGenesis initializes the epochs module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// set epoch info from genesis
	for _, epoch := range genState.Epochs {
		// Initialize empty epoch values via Cosmos SDK
		if epoch.StartTime.Equal(time.Time{}) || epoch.StartTime.IsZero() {
			epoch.StartTime = ctx.BlockTime()
		}

		epoch.CurrentEpochStartHeight = ctx.BlockHeight()

		k.SetEpochInfo(ctx, epoch)
	}
}

// ExportGenesis returns the epochs module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Epochs: k.AllEpochInfos(ctx),
	}
}
