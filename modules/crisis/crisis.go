package crisis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/irita-msg-parser/modules"
)

type crisisClient struct {
}

func NewClient() Client {
	return crisisClient{}
}

func (crisis crisisClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgVerifyInvariant).Type():
		docMsg := DocMsgVerifyInvariant{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
