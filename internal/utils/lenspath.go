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

// transactions
var transactionsLens = composeLenspath(dataLens, []string{"Transactions", "*"})
var vLens = composeLenspath(transactionsLens, []string{"v"})
var rLens = composeLenspath(transactionsLens, []string{"r"})
var sLens = composeLenspath(transactionsLens, []string{"s"})
var toLens = composeLenspath(transactionsLens, []string{"to"})
var fromLens = composeLenspath(transactionsLens, []string{"from"})

// blob tx
var blobTxSidecarLens = composeLenspath(transactionsLens, []string{"blobTxSidecar"})
var blobsLens = composeLenspath(blobTxSidecarLens, []string{"Blobs"})
var commitmentsLens = composeLenspath(blobTxSidecarLens, []string{"Commitments"})
var proofsLens = composeLenspath(blobTxSidecarLens, []string{"Proofs"})

// blob header
var blobFeeCapLens = composeLenspath(transactionsLens, []string{"blobFeeCap"})
var blobHashesLens = composeLenspath(transactionsLens, []string{"blobHashes"})
var blobGasLens = composeLenspath(transactionsLens, []string{"blobGas"})

// header
var headerLens = composeLenspath(dataLens, []string{"Header"})
var withdrawalsRootLens = composeLenspath(headerLens, []string{"withdrawalsRoot"})
var blobGasUsedLens = composeLenspath(headerLens, []string{"blobGasUsed"})
var excessBlobGasLens = composeLenspath(headerLens, []string{"excessBlobGas"})
var parentBeaconRootLens = composeLenspath(headerLens, []string{"parentBeaconBlockRoot"})

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
