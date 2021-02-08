package oracle

import (
	"github.com/irisnet/irismod/modules/oracle"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(oracle.AppModuleBasic{})
}
