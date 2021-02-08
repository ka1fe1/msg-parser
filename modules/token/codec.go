package token

import (
	"github.com/irisnet/irismod/modules/token"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
