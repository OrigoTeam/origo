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
package web3

import (
	"github.com/OrigoTeam/origo/version"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PublicAPI is the web3_ prefixed set of APIs in the Web3 JSON-RPC spec.
type PublicAPI struct{}

// NewPublicAPI creates an instance of the Web3 API.
func NewPublicAPI() *PublicAPI {
	return &PublicAPI{}
}

// ClientVersion returns the client version in the Web3 user agent format.
func (a *PublicAPI) ClientVersion() string {
	return version.Version()
}

// Sha3 returns the keccak-256 hash of the passed-in input.
func (a *PublicAPI) Sha3(input string) hexutil.Bytes {
	return crypto.Keccak256(hexutil.Bytes(input))
}
