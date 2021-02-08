package service

import (
	"github.com/irisnet/irismod/modules/service"
	"github.com/weichang-bianjie/irita-msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
