package utils

import (
	lensp "github.com/covalenthq/lenspath"
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"
)

// data
var dataLens = createLenspath([]string{"replicaEvent", "*", "data"})
var withdrawalsLens = composeLenspath(dataLens, []string{"Withdrawals"})
var uncleLens = composeLenspath(dataLens, []string{"Uncles"})
var blobTxSidecarsLens = composeLenspath(dataLens, []string{"BlobTxSidecars"})

// transactions
var transactionsLens = composeLenspath(dataLens, []string{"Transactions", "*"})
var vLens = composeLenspath(transactionsLens, []string{"v"})
var rLens = composeLenspath(transactionsLens, []string{"r"})
var sLens = composeLenspath(transactionsLens, []string{"s"})
var toLens = composeLenspath(transactionsLens, []string{"to"})
var fromLens = composeLenspath(transactionsLens, []string{"from"})

// blob data
var blobsLens = composeLenspath(blobTxSidecarsLens, []string{"Blobs"})
var commitmentsLens = composeLenspath(blobTxSidecarsLens, []string{"Commitments"})
var proofsLens = composeLenspath(blobTxSidecarsLens, []string{"Proofs"})

// blob tx fields
var blobFeeCapLens = composeLenspath(transactionsLens, []string{"blobFeeCap"})
var blobHashesLens = composeLenspath(transactionsLens, []string{"blobHashes"})
var blobGasLens = composeLenspath(transactionsLens, []string{"blobGas"})

// header
var headerLens = composeLenspath(dataLens, []string{"Header"})
var withdrawalsRootLens = composeLenspath(headerLens, []string{"withdrawalsRoot"})
var blobGasUsedLens = composeLenspath(headerLens, []string{"blobGasUsed"})
var excessBlobGasLens = composeLenspath(headerLens, []string{"excessBlobGas"})
var parentBeaconRootLens = composeLenspath(headerLens, []string{"parentBeaconBlockRoot"})

// EIP-7685 (EL-CL tx)
var requestsHashLens = composeLenspath(headerLens, []string{"requestsHash"})

// EIP-7702 (set EOA)
var dataTxLens = composeLenspath(transactionsLens, []string{"Data"})
var authListLens = composeLenspath(transactionsLens, []string{"AuthList"})
var chainIdLens = composeLenspath(authListLens, []string{"chainId"})
var addressLens = composeLenspath(authListLens, []string{"address"})
var nonceLens = composeLenspath(authListLens, []string{"nonce"})
var yParityLens = composeLenspath(authListLens, []string{"yParity"})
var rTxLens = composeLenspath(authListLens, []string{"r"})
var sTxLens = composeLenspath(authListLens, []string{"s"})

// utilities for lenspath

func unwrapType(data map[string]interface{}, lenspath *lensp.Lenspath, nonNilType string) {
	lenspathSetter(data, lenspath, func(leafd any) any {
		if leafd == nil {
			return nil
		}

		mp := leafd.(map[string]interface{})

		return mp[nonNilType]
	})
}

func wrapType(data map[string]interface{}, lenspath *lensp.Lenspath, nonNilType string) {
	lenspathSetter(data, lenspath, func(leafd any) any {
		wType := nonNilType
		if leafd == nil {
			wType = "null"
		}

		return goavro.Union(wType, leafd)
	})
}

func lenspathSetter(data map[string]interface{}, lenspath *lensp.Lenspath, setter func(leafd any) any) {
	err := lenspath.Setter(data, setter)

	if err != nil {
		log.Fatal(err)
	}
}

func composeLenspath(prevLenspath *lensp.Lenspath, lens []lensp.Lens) *lensp.Lenspath {
	lenspath, err := prevLenspath.Compose(lens)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	return lenspath
}

func createLenspath(lens []lensp.Lens) *lensp.Lenspath {
	lenspath, err := lensp.Create(lens)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	return lenspath
}
