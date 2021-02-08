package slashing

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/irita-msg-parser/modules"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
