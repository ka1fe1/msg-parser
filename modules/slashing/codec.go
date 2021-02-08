package slashing

import (
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
}
