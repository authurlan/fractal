// Copyright 2019 The Fractal Team Authors
// This file is part of the fractal project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package plugin

import (
	"math/big"

	"github.com/fractalplatform/fractal/common"
	"github.com/fractalplatform/fractal/state"
	"github.com/fractalplatform/fractal/types"
)

// IPM plugin manager interface.
type IPM interface {
	IAccount
	IAsset
	IConsensus
	IContract
	IFee
	ISinger
	ExecTx(arg interface{}) ([]byte, error)
}

// IAccount account manager interface.
type IAccount interface {
	GetNonce(account string) (uint64, error)
	SetNonce(account string, nonce uint64) error

	CreateAccount(pubKey common.PubKey, description string) ([]byte, error)

	GetCode(account string) ([]byte, error)
	SetCode(account string, code []byte) error

	GetBalance(account string, assetID uint64) (*big.Int, error)

	CanTransfer(account string, assetID uint64, value *big.Int) (bool, error)

	TransferAsset(from, to string, assetID uint64, value *big.Int) error
}

type IAsset interface {
	IssueAsset(account string, assetName string, symbol string, amount *big.Int,
		dec uint64, founder string, owner string, limit *big.Int, description string, asm IAsset) ([]byte, error)
}

type IConsensus interface {

	// VerifySeal checks whether the crypto seal on a header is valid according to the consensus rules of the given engine.
	VerifySeal(header *types.Header) error

	// Prepare initializes the consensus fields of a block header according to the rules of a particular engine. The changes are executed inline.
	Prepare(header *types.Header, txs []*types.Transaction, receipts []*types.Receipt, state *state.StateDB) error

	// Finalize assembles the final block.
	Finalize(header *types.Header, txs []*types.Transaction, receipts []*types.Receipt, state *state.StateDB) (*types.Block, error)
}

type IContract interface {
}

type IFee interface {
}

type ISinger interface {
	Sign(interface{}) ([]byte, error)
	Recover(signer types.Signer, tx *types.Transaction) error
}
