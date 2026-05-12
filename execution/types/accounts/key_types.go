// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"github.com/erigontech/erigon/common"
	"github.com/erigontech/erigon/common/empty"
)

type Address = common.Address

var ZeroAddress = common.Address{}
var NilAddress = common.Address{}

func InternAddress(a common.Address) common.Address {
	return a
}

type StorageKey = common.Hash

var ZeroKey = common.Hash{}
var NilKey = common.Hash{}

func InternKey(k common.Hash) common.Hash {
	return k
}

type CodeHash = common.Hash

var ZeroCodeHash = common.Hash{}
var NilCodeHash = common.Hash{}
var EmptyCodeHash = empty.CodeHash

func InternCodeHash(k common.Hash) common.Hash {
	return k
}
