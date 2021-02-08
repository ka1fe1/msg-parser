package auth

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/irita-msg-parser/modules"
)

type authClient struct {
}

func NewClient() Client {
	return authClient{}
}

func (auth authClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := false

	return msgDocInfo, ok
}
