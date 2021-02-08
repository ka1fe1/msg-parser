package evidence

import (
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(evidence.AppModuleBasic{})
}
