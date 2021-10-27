package types

import (
	"math/big"

	"github.com/ubiq/go-ubiq/common"
)

const (
	BloomByteLength = 256
	BloomBitLength  = 8 * BloomByteLength
)

type BlockResult struct {
	Hash            common.Hash
	TotalDifficulty *big.Int
	Header          *Header
	Transactions    []*Transaction
	Uncles          []*Header
	Receipts        []*Receipt
	Senders         []common.Address
}

type BlockSpecimen struct {
	Hash         common.Hash
	Header       *Header
	Transactions []*Transaction
	Uncles       []*Header
	State        *StateSpecimen
}

type StateSpecimen struct {
	AccountRead []*accountRead
	StorageRead []*storageRead
	CodeRead    []*codeRead
}

type BlockNonce [8]byte

type Bloom [BloomByteLength]byte

type Header struct {
	ParentHash  common.Hash    `json:"parentHash"`
	UncleHash   common.Hash    `json:"sha3Uncles"`
	Coinbase    common.Address `json:"miner"`
	Root        common.Hash    `json:"stateRoot"`
	TxHash      common.Hash    `json:"transactionsRoot"`
	ReceiptHash common.Hash    `json:"receiptsRoot"`
	Bloom       Bloom          `json:"logsBloom"`
	Difficulty  *big.Int       `json:"difficulty"`
	Number      *big.Int       `json:"number"`
	GasLimit    uint64         `json:"gasLimit"`
	GasUsed     uint64         `json:"gasUsed"`
	Time        uint64         `json:"timestamp"`
	Extra       []byte         `json:"extraData"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`
	BaseFee     *big.Int       `json:"baseFeePerGas" rlp:"-"`
}

type Transaction struct {
	AccountNonce uint64          `json:"nonce"    `
	Price        *big.Int        `json:"gasPrice" `
	GasLimit     uint64          `json:"gas"      `
	Sender       common.Address  `json:"from"     `
	Recipient    *common.Address `json:"to,omitempty"  rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    `
	Payload      []byte          `json:"input"    `
}

type Logs struct {
	Address     common.Address `json:"address" `
	Topics      []common.Hash  `json:"topics" `
	Data        []byte         `json:"data" `
	BlockNumber uint64         `json:"blockNumber"`
	TxHash      common.Hash    `json:"transactionHash" `
	TxIndex     uint           `json:"transactionIndex"`
	BlockHash   common.Hash    `json:"blockHash"`
	Index       uint           `json:"logIndex"`
	Removed     bool           `json:"removed"`
}

type Receipt struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	TxHash            common.Hash
	ContractAddress   common.Address
	Logs              []*Logs
	GasUsed           uint64
}

type accountRead struct {
	Address  common.Address
	Nonce    uint64
	Balance  *big.Int
	CodeHash common.Hash
}

type storageRead struct {
	Account common.Address
	SlotKey common.Hash
	Value   common.Hash
}

type codeRead struct {
	Hash common.Hash
	Code []byte
}
