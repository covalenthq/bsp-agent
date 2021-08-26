package types

import (
	"math/big"

	"github.com/ubiq/go-ubiq/common"
)

type BlockSpecimen struct {
	Address  common.Address
	Nonce    uint64
	Balance  *big.Int
	CodeHash common.Hash
	Account  common.Address
	SlotKey  common.Hash
	Value    common.Hash
	Hash     common.Hash
	Code     []byte
}
