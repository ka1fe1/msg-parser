package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

type (
	TxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}

	SdkMsg    types.Msg
	LegacyMsg legacytx.LegacyMsg
)

func MsgTypeURL(msg SdkMsg) string {
	return types.MsgTypeURL(msg)
}
