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

package types

import (
	"github.com/ethereum/go-ethereum/common"
	evmostypes "github.com/OrigoTeam/origo/types"
)

// NewGasMeter returns an instance of GasMeter
func NewGasMeter(
	contract common.Address,
	participant common.Address,
	cumulativeGas uint64,
) GasMeter {
	return GasMeter{
		Contract:      contract.String(),
		Participant:   participant.String(),
		CumulativeGas: cumulativeGas,
	}
}

// Validate performs a stateless validation of a Incentive
func (gm GasMeter) Validate() error {
	if err := evmostypes.ValidateAddress(gm.Contract); err != nil {
		return err
	}

	return evmostypes.ValidateAddress(gm.Participant)
}
