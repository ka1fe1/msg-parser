package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	cdc "github.com/kaifei-bianjie/msg-parser/codec"
	models "github.com/kaifei-bianjie/msg-parser/types"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func CreateMsgDocInfo(msg sdk.Msg, handler func() (Msg, []string)) MsgDocInfo {
	var (
		docTxMsg models.TxMsg
		signers  []string
		addrs    []string
	)

	m, addrcollections := handler()

	m.BuildMsg(msg)
	docTxMsg = models.TxMsg{
		Type: m.GetType(),
		Msg:  m,
	}

	_, signers = models.BuildDocSigners(msg.GetSigners())
	addrs = append(addrs, signers...)
	addrs = append(addrs, addrcollections...)

	return MsgDocInfo{
		DocTxMsg: docTxMsg,
		Signers:  signers,
		Addrs:    addrs,
	}
}

func ConvertMsg(v proto.Message, msg proto.Message) {
	bytes, _ := proto.Marshal(v)
	_ = proto.Unmarshal(bytes, msg)
}

func ConvertAny(v *types.Any) string {
	if v != nil {
		return utils.MarshalJsonIgnoreErr(v)
	}
	return ""
}

func UnmarshalTibcAcknowledgement(bytesdata []byte) string {
	var result TIBCAcknowledgement
	if err := cdc.GetMarshaler().Unmarshal(bytesdata, &result); err == nil {
		return result.String()
	}
	return ""
}
