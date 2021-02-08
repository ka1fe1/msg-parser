package bank

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(bank.AppModuleBasic{})
}
