package identity

import (
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(identity.AppModuleBasic{})
}
