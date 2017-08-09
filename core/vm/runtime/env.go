// Copyright 2015 The cidas Authors
// This file is part of the cidas library.
//
// The cidas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cidas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cidas library. If not, see <http://www.gnu.org/licenses/>.

package runtime

import (
	"math/big"

	"github.com/cidas/cidas/common"
	"github.com/cidas/cidas/core"
	"github.com/cidas/cidas/core/state"
	"github.com/cidas/cidas/core/vm"
)

func NewEnv(cfg *Config, state *state.StateDB) *vm.EVM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:      cfg.Origin,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    new(big.Int).SetUint64(cfg.GasLimit),
		GasPrice:    new(big.Int),
	}

	return vm.NewEVM(context, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}
