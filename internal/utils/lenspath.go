package utils

import (
	"github.com/linkedin/goavro/v2"
	log "github.com/sirupsen/logrus"
	lensp "github.com/covalenthq/lenspath"
)

var dataLens = createLenspath([]string{"replicaEvent", "*", "data"})
var transactionsLens = composeLenspath(dataLens, []string{"Transactions", "*"})
var vLens = composeLenspath(transactionsLens, []string{"v"})
var rLens = composeLenspath(transactionsLens, []string{"r"})
var sLens = composeLenspath(transactionsLens, []string{"s"})
var toLens = composeLenspath(transactionsLens, []string{"to"})
var fromLens = composeLenspath(transactionsLens, []string{"from"})

var headerLens = composeLenspath(dataLens, []string{"Header"})
var withdrawalsRootLens = composeLenspath(headerLens, []string{"withdrawalsRoot"})

var withdrawalsLens = composeLenspath(dataLens, []string{"Withdrawals"})
var uncleLens = composeLenspath(dataLens, []string{"Uncles"})

// utilities for lenspath

func unwrapType(data map[string]interface{}, lenspath *lensp.Lenspath, nonNilType string) map[string]interface{} {
	return lenspathSetter(data, lenspath, func(leafd any) any {
		if leafd == nil {
			return nil
		}

		mp := leafd.(map[string]interface{})
		return mp[nonNilType]
	})
}

func wrapType(data map[string]interface{}, lenspath *lensp.Lenspath, nonNilType string) map[string]interface{} {
	return lenspathSetter(data, lenspath, func(leafd any) any {
		wType := nonNilType
		if leafd == nil {
			wType = "null"
		}

		return goavro.Union(wType, leafd)
	})
}

func lenspathSetter(data map[string]interface{}, lenspath *lensp.Lenspath, setter func(leafd any) any) map[string]interface{} {
	err := lenspath.Setter(data, setter)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func composeLenspath(prevLenspath *lensp.Lenspath, lens []lensp.Lens) *lensp.Lenspath {
	if lenspath, err := prevLenspath.Compose(lens); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return lenspath
	}
}

func createLenspath(lens []lensp.Lens) *lensp.Lenspath {
	if lenspath, err := lensp.Create(lens); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return lenspath
	}
}
