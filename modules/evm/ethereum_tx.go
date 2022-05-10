package evm

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
	"github.com/kaifei-bianjie/msg-parser/utils/contracts"
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
	Ex   struct {
		DdcType   string `bson:"ddc_type"`
		DdcId     int64  `bson:"ddc_id"`
		DdcMethod string `bson:"ddc_method"`
		DdcInputs string `bson:"ddc_inputs"`
		To        string `bson:"to"`
	} `bson:"ex"`
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
		if txData.GetTo() != nil {
			doctx.Ex.To = txData.GetTo().String()
		}
		doctx.Data = utils.MarshalJsonIgnoreErr(txData)
		inputDataStr := hex.EncodeToString(txData.GetData())
		if err := parseContractsInput(inputDataStr, doctx); err != nil {
			fmt.Println(err.Error())
		}
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

func parseContractsInput(inputDataStr string, doctx *DocMsgEthereumTx) error {
	ddcMethodId := inputDataStr[:8]
	methodMap, err := contracts.GetDDCSupportMethod()
	if err != nil {
		return err
	}
	if val, ok := methodMap[ddcMethodId]; ok {
		doctx.Ex.DdcType = val.Contracts
		doctx.Ex.DdcMethod = val.Method.Name
		inputData, err := hex.DecodeString(inputDataStr[8:])
		if err != nil {
			return err
		}

		inputs, err := val.Method.Inputs.Unpack(inputData)
		if err != nil {
			return err
		}
		doctx.Ex.DdcInputs = utils.MarshalJsonIgnoreErr(inputs)
	}

	return nil
}
