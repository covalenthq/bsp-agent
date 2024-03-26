//nolint:stylecheck,revive
package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ubiq/go-ubiq/common"
)

const (
	BloomByteLength = 256
	BloomBitLength  = 8 * BloomByteLength
)

type BlockReplica struct {
	Type            string
	NetworkId       uint64
	Hash            common.Hash
	TotalDifficulty *BigInt
	Header          *Header
	Transactions    []*Transaction
	Uncles          []*Header `json:"uncles,omitempty"`
	Receipts        []*Receipt
	Senders         []common.Address
	State           *StateSpecimen `json:"State"`
	Withdrawals     []*Withdrawal
}
type StateSpecimen struct {
	AccountRead   []*AccountRead
	StorageRead   []*StorageRead
	CodeRead      []*CodeRead
	BlockhashRead []*BlockhashRead
}

type BlockhashRead struct {
	BlockNumber uint64
	BlockHash   common.Hash
}

type Withdrawal struct {
	Index     uint64         `json:"index"`          // monotonically increasing identifier issued by consensus layer
	Validator uint64         `json:"validatorIndex"` // index of validator associated with withdrawal
	Address   common.Address `json:"address"`        // target address for withdrawn ether
	Amount    uint64         `json:"amount"`         // value of withdrawal in Gwei
}

type BlockNonce [8]byte

type Bloom [BloomByteLength]byte

type Header struct {
	ParentHash       common.Hash    `json:"parentHash"`
	UncleHash        common.Hash    `json:"sha3Uncles"`
	Coinbase         common.Address `json:"miner"`
	Root             common.Hash    `json:"stateRoot"`
	TxHash           common.Hash    `json:"transactionsRoot"`
	ReceiptHash      common.Hash    `json:"receiptsRoot"`
	Bloom            Bloom          `json:"logsBloom"`
	Difficulty       *BigInt        `json:"difficulty"`
	Number           *BigInt        `json:"number"`
	GasLimit         uint64         `json:"gasLimit"`
	GasUsed          uint64         `json:"gasUsed"`
	Time             uint64         `json:"timestamp"`
	Extra            []byte         `json:"extraData"`
	MixDigest        common.Hash    `json:"mixHash"`
	Nonce            BlockNonce     `json:"nonce"`
	BaseFee          *BigInt        `json:"baseFeePerGas"`
	WithdrawalsHash  *common.Hash   `json:"withdrawalsRoot" rlp:"nil,optional"`
	BlobGasUsed      *uint64        `json:"blobGasUsed" rlp:"optional"`
	ExcessBlobGas    *uint64        `json:"excessBlobGas" rlp:"optional"`
	ParentBeaconRoot *common.Hash   `json:"parentBeaconBlockRoot" rlp:"optional"`
}

type Transaction struct {
	Type          byte            `json:"type"`
	AccessList    AccessList      `json:"accessList"`
	ChainId       *BigInt         `json:"chainId"`
	AccountNonce  uint64          `json:"nonce"`
	Price         *BigInt         `json:"gasPrice"`
	GasLimit      uint64          `json:"gas"`
	GasTipCap     *BigInt         `json:"gasTipCap"`
	GasFeeCap     *BigInt         `json:"gasFeeCap"`
	Sender        *common.Address `json:"from" rlp:"nil"`
	Recipient     *common.Address `json:"to" rlp:"nil"` // nil means contract creation
	Amount        *BigInt         `json:"value"`
	Payload       []byte          `json:"input"`
	V             *BigInt         `json:"v" rlp:"nilString"`
	R             *BigInt         `json:"r" rlp:"nilString"`
	S             *BigInt         `json:"s" rlp:"nilString"`
	BlobFeeCap    *big.Int        `json:"blobFeeCap" rlp:"optional"`
	BlobHashes    []common.Hash   `json:"blobHashes" rlp:"optional"`
	BlobGas       uint64          `json:"blobGas" rlp:"optional"`
	BlobTxSidecar *BlobTxSidecar  `json:"blobTxSidecar" rlp:"optional"`
}

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     common.Address `json:"address"        gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys"    gencodec:"required"`
}

type Logs struct {
	Address     common.Address `json:"address"`
	Topics      []common.Hash  `json:"topics"`
	Data        []byte         `json:"data"`
	BlockNumber uint64         `json:"blockNumber"`
	TxHash      common.Hash    `json:"transactionHash"`
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

type AccountRead struct {
	Address  common.Address
	Nonce    uint64
	Balance  *BigInt
	CodeHash common.Hash
}

type StorageRead struct {
	Account common.Address
	SlotKey common.Hash
	Value   common.Hash
}

type CodeRead struct {
	Hash common.Hash
	Code []byte
}

// BlobTxSidecar contains the blobs of a blob transaction.
type BlobTxSidecar struct {
	Blobs       []kzg4844.Blob       // Blobs needed by the blob pool
	Commitments []kzg4844.Commitment // Commitments needed by the blob pool
	Proofs      []kzg4844.Proof      // Proofs needed by the blob pool
}
