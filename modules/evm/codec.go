package evm

import (
	"github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/tharsis/ethermint/x/evm"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
