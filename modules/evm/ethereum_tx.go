package evm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
	evm "github.com/tharsis/ethermint/x/evm/types"
)

// MsgEthereumTx encapsulates an Ethereum transaction as an SDK message.
type DocMsgEthereumTx struct {
	// inner transaction data
	Data string `bson:"data"`
	// encoded storage size of the transaction
	Size_ float64 `bson:"size"`
	// transaction hash in hex format
	Hash string `bson:"hash"`
	// ethereum signer address in hex format. This address value is checked
	// against the address derived from the signature (V, R, S) using the
	// secp256k1 elliptic curve
	From string `bson:"from"`
}

func (doctx *DocMsgEthereumTx) GetType() string {
	return MsgTypeEthereumTx
}

func (doctx *DocMsgEthereumTx) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgEthereumTx)
	doctx.Size_ = msg.Size_
	doctx.Hash = msg.Hash
	doctx.From = msg.From
	if txData, err := evm.UnpackTxData(msg.Data); err == nil {
		doctx.Data = utils.MarshalJsonIgnoreErr(txData)
	}
}

func (m *DocMsgEthereumTx) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
	)
	msg, ok := v.(*MsgEthereumTx)
	if ok {
		addrs = append(addrs, msg.From)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
