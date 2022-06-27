package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestFeegrant() {
	cases := []SubTest{
		{
			"GrantAllowance",
			GrantAllowance,
		},
		{
			"RevokeAllowance",
			RevokeAllowance,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func GrantAllowance(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0ab5010ab2010a2a2f636f736d6f732e6665656772616e742e763162657461312e4d73674772616e74416c6c6f77616e63651283010a2a696161317974656d7a32787171327337337574337973386d6364367a63613235363461356c6668746d33122a6961613130366c6367356d38683363646177756e376332727277706a3771336e636672396b33777877781a290a272f636f736d6f732e6665656772616e742e763162657461312e4261736963416c6c6f77616e636512640a4c0a400a192f636f736d6f732e63727970746f2e736d322e5075624b657912230a210274ccc6501f48c739ac5997021d6349b5fc7f106dd1092fbbc849b36438035bdb12040a02080118d1f10d12140a0e0a0475676173120632303030303010c09a0c1a40310f7191403d9cc6be07d5d4d5226c278ac5553842cd54abbc980842eeb4e262a19ff8d11faa1e26c785e4aef3fa9120ac10775a265724fd866238b11bbb7e21")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if feegrantDoc, ok := s.Feegrant.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(feegrantDoc))
		}
	}
}

func RevokeAllowance(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0a8a010a87010a2b2f636f736d6f732e6665656772616e742e763162657461312e4d73675265766f6b65416c6c6f77616e636512580a2a696161317974656d7a32787171327337337574337973386d6364367a63613235363461356c6668746d33122a6961613130366c6367356d38683363646177756e376332727277706a3771336e636672396b337778777812640a4c0a400a192f636f736d6f732e63727970746f2e736d322e5075624b657912230a210274ccc6501f48c739ac5997021d6349b5fc7f106dd1092fbbc849b36438035bdb12040a02080118e4ef0d12140a0e0a0475676173120632303030303010c09a0c1a40d9cbab0bc6ab21b75a5e78dd3074c03cb0407d2911352aa8ef1d447e2cc3f1a7c209ab780755ec3daead7fa4ba853a2234bf80b9bfd68759083ff4b9198146e3")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if feegrantDoc, ok := s.Feegrant.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(feegrantDoc))
		}
	}
}
