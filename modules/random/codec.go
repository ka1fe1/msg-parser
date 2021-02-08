package random

import (
	"github.com/irisnet/irismod/modules/random"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(random.AppModuleBasic{})
}
