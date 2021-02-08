package nft

import (
	"github.com/irisnet/irismod/modules/nft"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
