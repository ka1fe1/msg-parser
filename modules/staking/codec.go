package staking

import (
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(staking.AppModuleBasic{})
}
