package distribution

import (
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(distribution.AppModuleBasic{})
}
