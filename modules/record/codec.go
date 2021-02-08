package record

import (
	"github.com/irisnet/irismod/modules/record"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
