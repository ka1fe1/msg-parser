package params

import (
	"github.com/bianjieai/iritamod/modules/params"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
