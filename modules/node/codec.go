package node

import (
	"github.com/kaifei-bianjie/msg-parser/codec"
	//"github.com/bianjieai/spartan-cosmos/module/node"
	node "github.com/bianjieai/spartan-cosmos/module/node/module"
)

func init() {
	codec.RegisterAppModules(
		node.AppModuleBasic{},
	)
}
