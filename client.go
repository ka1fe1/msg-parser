package msg_sdk

import (
	"github.com/weichang-bianjie/irita-msg-parser/codec"
	"github.com/weichang-bianjie/irita-msg-parser/modules/auth"
	"github.com/weichang-bianjie/irita-msg-parser/modules/bank"
	"github.com/weichang-bianjie/irita-msg-parser/modules/coinswap"
	"github.com/weichang-bianjie/irita-msg-parser/modules/crisis"
	"github.com/weichang-bianjie/irita-msg-parser/modules/distribution"
	"github.com/weichang-bianjie/irita-msg-parser/modules/evidence"
	"github.com/weichang-bianjie/irita-msg-parser/modules/gov"
	"github.com/weichang-bianjie/irita-msg-parser/modules/htlc"
	"github.com/weichang-bianjie/irita-msg-parser/modules/ibc"
	"github.com/weichang-bianjie/irita-msg-parser/modules/identity"
	"github.com/weichang-bianjie/irita-msg-parser/modules/nft"
	"github.com/weichang-bianjie/irita-msg-parser/modules/oracle"
	"github.com/weichang-bianjie/irita-msg-parser/modules/params"
	"github.com/weichang-bianjie/irita-msg-parser/modules/random"
	"github.com/weichang-bianjie/irita-msg-parser/modules/record"
	"github.com/weichang-bianjie/irita-msg-parser/modules/service"
	"github.com/weichang-bianjie/irita-msg-parser/modules/slashing"
	"github.com/weichang-bianjie/irita-msg-parser/modules/staking"
	"github.com/weichang-bianjie/irita-msg-parser/modules/token"
	"github.com/weichang-bianjie/irita-msg-parser/modules/upgrade"
)

type MsgClient struct {
	Auth         auth.Client
	Bank         bank.Client
	Staking      staking.Client
	Crisis       crisis.Client
	Distribution distribution.Client
	Evidence     evidence.Client
	Gov          gov.Client
	Ibc          ibc.Client
	Params       params.Client
	Slashing     slashing.Client
	Upgrade      upgrade.Client
	Service      service.Client
	Nft          nft.Client
	Token        token.Client
	Random       random.Client
	Oracle       oracle.Client
	Htlc         htlc.Client
	Record       record.Client
	Coinswap     coinswap.Client
	Identity     identity.Client
	//Wasm         wasm.Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Auth:         auth.NewClient(),
		Bank:         bank.NewClient(),
		Crisis:       crisis.NewClient(),
		Distribution: distribution.NewClient(),
		Evidence:     distribution.NewClient(),
		Gov:          gov.NewClient(),
		Ibc:          ibc.NewClient(),
		Params:       params.NewClient(),
		Slashing:     slashing.NewClient(),
		Upgrade:      upgrade.NewClient(),
		Staking:      staking.NewClient(),
		Service:      service.NewClient(),
		Nft:          nft.NewClient(),
		Record:       record.NewClient(),
		Random:       random.NewClient(),
		Oracle:       oracle.NewClient(),
		Htlc:         htlc.NewClient(),
		Token:        token.NewClient(),
		Coinswap:     coinswap.NewClient(),
		Identity:     identity.NewClient(),
		//Wasm:         wasm.NewClient(),
	}
}
