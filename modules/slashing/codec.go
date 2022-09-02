package slashing

import (
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
	"github.com/cosmos/cosmos-sdk/x/slashing"

	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
	codec.RegisterAppModules(iritaslashing.AppModuleBasic{})
}
