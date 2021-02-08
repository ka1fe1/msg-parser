package crisis

import (
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(crisis.AppModuleBasic{})
}
