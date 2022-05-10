package contracts

import (
	"encoding/hex"
	ABI "github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

type MethodData struct {
	Method    ABI.Method
	Contracts string
}

var ddcABIsMap = map[string]string{
	"DDC1155":   _DDC1155ABI,
	"DDC721":    _DDC721ABI,
	"Authority": _AuthorityABI,
	"Charge":    _ChargeABI,
}

func GetDDCSupportMethod() (map[string]MethodData, error) {
	methodMap := make(map[string]MethodData)
	for key, ddcABI := range ddcABIsMap {
		abiServe, err := ABI.JSON(strings.NewReader(ddcABI))
		if err != nil {
			return nil, err
		}
		for _, method := range abiServe.Methods {
			methodMap[hex.EncodeToString(method.ID)] = MethodData{
				Method:    method,
				Contracts: key,
			}
		}
	}

	return methodMap, nil
}
