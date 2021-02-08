package upgrade

import (
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
