package htlc

import (
	"github.com/irisnet/irismod/modules/htlc"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(htlc.AppModuleBasic{})
}
