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

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/OrigoTeam/origo/x/inflation/types"
)

// GetEpochIdentifier gets the epoch identifier
func (k Keeper) GetEpochIdentifier(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEpochIdentifier)
	if len(bz) == 0 {
		return ""
	}

	return string(bz)
}

// SetEpochsPerPeriod stores the epoch identifier
func (k Keeper) SetEpochIdentifier(ctx sdk.Context, epochIdentifier string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEpochIdentifier, []byte(epochIdentifier))
}

// GetEpochsPerPeriod gets the epochs per period
func (k Keeper) GetEpochsPerPeriod(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEpochsPerPeriod)
	if len(bz) == 0 {
		return 0
	}

	return int64(sdk.BigEndianToUint64(bz))
}

// SetEpochsPerPeriod stores the epochs per period
func (k Keeper) SetEpochsPerPeriod(ctx sdk.Context, epochsPerPeriod int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEpochsPerPeriod, sdk.Uint64ToBigEndian(uint64(epochsPerPeriod)))
}

// GetSkippedEpochs gets the number of skipped epochs
func (k Keeper) GetSkippedEpochs(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixSkippedEpochs)
	if len(bz) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetSkippedEpochs stores the number of skipped epochs
func (k Keeper) SetSkippedEpochs(ctx sdk.Context, skippedEpochs uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixSkippedEpochs, sdk.Uint64ToBigEndian(skippedEpochs))
}
