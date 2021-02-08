package ibc

import (
	"github.com/weichang-bianjie/irita-msg-parser/codec"
	ibcrecord "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record"
	ibc "gitlab.bianjie.ai/cschain/cschain/modules/ibc/core"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibcrecord.AppModuleBasic{})
}
