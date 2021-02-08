package coinswap

import (
	"github.com/irisnet/irismod/modules/coinswap"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(coinswap.AppModuleBasic{})
}
