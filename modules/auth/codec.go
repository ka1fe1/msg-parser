package auth

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
