package ibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type ibcClient struct {
}

func NewClient() Client {
	return ibcClient{}
}

func (ibc ibcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgRecvPacket:
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgCreateClient:
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateClient:
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	default:
		ok = false
	}
	return msgDocInfo, ok
}
