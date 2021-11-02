package msgs

import (
	tibctranfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibcpacket "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	//"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfer "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/modules/core/02-client/types"
	ibcconnect "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	ibc "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibcchannel "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	coinswap "github.com/irisnet/irismod/modules/coinswap/types"
	farm "github.com/irisnet/irismod/modules/farm/types"
	htlc "github.com/irisnet/irismod/modules/htlc/types"
	nft "github.com/irisnet/irismod/modules/nft/types"
	oracle "github.com/irisnet/irismod/modules/oracle/types"
	random "github.com/irisnet/irismod/modules/random/types"
	record "github.com/irisnet/irismod/modules/record/types"
	service "github.com/irisnet/irismod/modules/service/types"
	token "github.com/irisnet/irismod/modules/token/types"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

const (
	MsgTypeSend          = "send"
	MsgTypeMultiSend     = "multisend"
	MsgTypeNFTMint       = "mint_nft"
	MsgTypeNFTEdit       = "edit_nft"
	MsgTypeNFTTransfer   = "transfer_nft"
	MsgTypeTransferDenom = "transfer_denom"
	MsgTypeNFTBurn       = "burn_nft"
	MsgTypeIssueDenom    = "issue_denom"
	MsgTypeRecordCreate  = "create_record"

	MsgTypeCreatePool  = "create_pool"
	MsgTypeDestroyPool = "destroy_pool"
	MsgTypeAdjustPool  = "adjust_pool"
	MsgTypeStake       = "stake"
	MsgTypeUnstake     = "unstake"
	MsgTypeHarvest     = "harvest"

	MsgTypeMintToken          = "mint_token"
	MsgTypeBurnToken          = "burn_token"
	MsgTypeEditToken          = "edit_token"
	MsgTypeIssueToken         = "issue_token"
	MsgTypeTransferTokenOwner = "transfer_token_owner"

	MsgTypeDefineService             = "define_service"               // type for MsgDefineService
	MsgTypeBindService               = "bind_service"                 // type for MsgBindService
	MsgTypeUpdateServiceBinding      = "update_service_binding"       // type for MsgUpdateServiceBinding
	MsgTypeServiceSetWithdrawAddress = "service/set_withdraw_address" // type for MsgSetWithdrawAddress
	MsgTypeDisableServiceBinding     = "disable_service_binding"      // type for MsgDisableServiceBinding
	MsgTypeEnableServiceBinding      = "enable_service_binding"       // type for MsgEnableServiceBinding
	MsgTypeRefundServiceDeposit      = "refund_service_deposit"       // type for MsgRefundServiceDeposit
	MsgTypeCallService               = "call_service"                 // type for MsgCallService
	MsgTypeRespondService            = "respond_service"              // type for MsgRespondService
	MsgTypePauseRequestContext       = "pause_request_context"        // type for MsgPauseRequestContext
	MsgTypeStartRequestContext       = "start_request_context"        // type for MsgStartRequestContext
	MsgTypeKillRequestContext        = "kill_request_context"         // type for MsgKillRequestContext
	MsgTypeUpdateRequestContext      = "update_request_context"       // type for MsgUpdateRequestContext
	MsgTypeWithdrawEarnedFees        = "withdraw_earned_fees"         // type for MsgWithdrawEarnedFees

	MsgTypeStakeCreateValidator           = "create_validator"
	MsgTypeStakeEditValidator             = "edit_validator"
	MsgTypeStakeDelegate                  = "delegate"
	MsgTypeStakeBeginUnbonding            = "begin_unbonding"
	MsgTypeBeginRedelegate                = "begin_redelegate"
	MsgTypeUnjail                         = "unjail"
	MsgTypeSetWithdrawAddress             = "set_withdraw_address"
	MsgTypeWithdrawDelegatorReward        = "withdraw_delegator_reward"
	MsgTypeMsgFundCommunityPool           = "fund_community_pool"
	MsgTypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	MsgTypeSubmitProposal                 = "submit_proposal"
	MsgTypeDeposit                        = "deposit"
	MsgTypeVote                           = "vote"

	MsgTypeCreateHTLC = "create_htlc"
	MsgTypeClaimHTLC  = "claim_htlc"
	MsgTypeRefundHTLC = "refund_htlc"

	MsgTypeAddLiquidity    = "add_liquidity"
	MsgTypeRemoveLiquidity = "remove_liquidity"
	MsgTypeSwapOrder       = "swap_order"

	MsgTypeSubmitEvidence  = "submit_evidence"
	MsgTypeVerifyInvariant = "verify_invariant"

	//ibc client
	MsgTypeCreateClient             = "create_client"
	MsgTypeUpdateClient             = "update_client"
	MsgTypeUpgradeClient            = "upgrade_client"
	MsgTypeSubmitMisbehaviourClient = "submit_misbehaviour"

	//ibc connect
	MsgTypeConnectionOpenInit    = "connection_open_init"
	MsgTypeConnectionOpenTry     = "connection_open_try"
	MsgTypeConnectionOpenAck     = "connection_open_ack"
	MsgTypeConnectionOpenConfirm = "connection_open_confirm"

	//ibc channel
	MsgTypeChannelOpenInit     = "channel_open_init"
	MsgTypeChannelOpenTry      = "channel_open_try"
	MsgTypeChannelOpenAck      = "channel_open_ack"
	MsgTypeChannelOpenConfirm  = "channel_open_confirm"
	MsgTypeChannelCloseInit    = "channel_close_init"
	MsgTypeChannelCloseConfirm = "channel_close_confirm"
	MsgTypeTimeout             = "timeout_packet"
	MsgTypeTimeoutOnClose      = "timeout_on_close_packet"
	MsgTypeAcknowledgement     = "acknowledge_packet"

	MsgTypeRecvPacket  = "recv_packet"
	MsgTypeIBCTransfer = "transfer"

	TxTypeRequestRand = "request_rand"

	TxTypeCreateFeed = "create_feed"
	TxTypeEditFeed   = "edit_feed"
	TxTypePauseFeed  = "pause_feed"
	TxTypeStartFeed  = "start_feed"

	MsgTypeUpdateAdmin         = "update_contract_admin"
	MsgTypeClearAdmin          = "clear_contract_admin"
	MsgTypeExecuteContract     = "execute"
	MsgTypeInstantiateContract = "instantiate"
	MsgTypeMigrateContract     = "migrate"
	MsgTypeStoreCode           = "store_code"

	MsgTypeTIBCNftTransfer     = "tibc_nft_transfer"
	MsgTypeTIBCRecvPacket      = "tibc_recv_packet"
	MsgTypeTIBCUpdateClient    = "tibc_update_client"
	MsgTypeTIBCAcknowledgement = "tibc_acknowledge_packet"
	MsgTypeTIBCCleanPacket     = "clean_packet"
	MsgTypeTIBCRecvCleanPacket = "recv_clean_packet"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin

	MsgSend      = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend

	MsgNFTMint       = nft.MsgMintNFT
	MsgNFTEdit       = nft.MsgEditNFT
	MsgNFTTransfer   = nft.MsgTransferNFT
	MsgNFTBurn       = nft.MsgBurnNFT
	MsgIssueDenom    = nft.MsgIssueDenom
	MsgTransferDenom = nft.MsgTransferDenom

	MsgUnstake     = farm.MsgUnstake
	MsgStake       = farm.MsgStake
	MsgCreatePool  = farm.MsgCreatePool
	MsgDestroyPool = farm.MsgDestroyPool
	MsgAdjustPool  = farm.MsgAdjustPool
	MsgHarvest     = farm.MsgHarvest

	MsgDefineService  = service.MsgDefineService
	MsgBindService    = service.MsgBindService
	MsgCallService    = service.MsgCallService
	MsgRespondService = service.MsgRespondService

	MsgUpdateServiceBinding  = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress    = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding  = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit  = service.MsgRefundServiceDeposit
	MsgPauseRequestContext   = service.MsgPauseRequestContext
	MsgStartRequestContext   = service.MsgStartRequestContext
	MsgKillRequestContext    = service.MsgKillRequestContext
	MsgUpdateRequestContext  = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees    = service.MsgWithdrawEarnedFees

	MsgRecordCreate = record.MsgCreateRecord

	MsgIssueToken         = token.MsgIssueToken
	MsgEditToken          = token.MsgEditToken
	MsgBurnToken          = token.MsgBurnToken
	MsgMintToken          = token.MsgMintToken
	MsgTransferTokenOwner = token.MsgTransferTokenOwner

	MsgStakeCreate                 = stake.MsgCreateValidator
	MsgStakeEdit                   = stake.MsgEditValidator
	MsgStakeDelegate               = stake.MsgDelegate
	MsgStakeBeginUnbonding         = stake.MsgUndelegate
	MsgBeginRedelegate             = stake.MsgBeginRedelegate
	MsgUnjail                      = slashing.MsgUnjail
	MsgStakeSetWithdrawAddress     = distribution.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward     = distribution.MsgWithdrawDelegatorReward
	MsgFundCommunityPool           = distribution.MsgFundCommunityPool
	MsgWithdrawValidatorCommission = distribution.MsgWithdrawValidatorCommission
	StakeValidator                 = stake.Validator
	Delegation                     = stake.Delegation
	UnbondingDelegation            = stake.UnbondingDelegation

	MsgDeposit        = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	TextProposal      = gov.TextProposal
	MsgVote           = gov.MsgVote
	Proposal          = gov.Proposal
	SdkVote           = gov.Vote
	GovContent        = gov.Content

	MsgSwapOrder       = coinswap.MsgSwapOrder
	MsgAddLiquidity    = coinswap.MsgAddLiquidity
	MsgRemoveLiquidity = coinswap.MsgRemoveLiquidity

	MsgClaimHTLC  = htlc.MsgClaimHTLC
	MsgCreateHTLC = htlc.MsgCreateHTLC

	MsgSubmitEvidence  = evidence.MsgSubmitEvidence
	MsgVerifyInvariant = crisis.MsgVerifyInvariant

	FungibleTokenPacketData = ibctransfer.FungibleTokenPacketData
	MsgRecvPacket           = ibc.MsgRecvPacket
	MsgTransfer             = ibctransfer.MsgTransfer

	MsgCreateClient       = ibcclient.MsgCreateClient
	MsgUpdateClient       = ibcclient.MsgUpdateClient
	MsgSubmitMisbehaviour = ibcclient.MsgSubmitMisbehaviour
	MsgUpgradeClient      = ibcclient.MsgUpgradeClient

	MsgConnectionOpenInit    = ibcconnect.MsgConnectionOpenInit
	MsgConnectionOpenAck     = ibcconnect.MsgConnectionOpenAck
	MsgConnectionOpenConfirm = ibcconnect.MsgConnectionOpenConfirm
	MsgConnectionOpenTry     = ibcconnect.MsgConnectionOpenTry

	Acknowledgement        = ibc.Acknowledgement
	MsgChannelOpenInit     = ibcchannel.MsgChannelOpenInit
	MsgChannelOpenTry      = ibcchannel.MsgChannelOpenTry
	MsgChannelOpenAck      = ibcchannel.MsgChannelOpenAck
	MsgChannelOpenConfirm  = ibcchannel.MsgChannelOpenConfirm
	MsgChannelCloseConfirm = ibcchannel.MsgChannelCloseConfirm
	MsgChannelCloseInit    = ibcchannel.MsgChannelCloseInit
	MsgAcknowledgement     = ibcchannel.MsgAcknowledgement
	MsgTimeout             = ibcchannel.MsgTimeout
	MsgTimeoutOnClose      = ibcchannel.MsgTimeoutOnClose

	MsgCreateFeed = oracle.MsgCreateFeed
	MsgEditFeed   = oracle.MsgEditFeed
	MsgPauseFeed  = oracle.MsgPauseFeed
	MsgStartFeed  = oracle.MsgStartFeed

	MsgRequestRandom = random.MsgRequestRandom

	//MsgStoreCode           = wasm.MsgStoreCode
	//MsgInstantiateContract = wasm.MsgInstantiateContract
	//MsgExecuteContract     = wasm.MsgExecuteContract
	//MsgMigrateContract     = wasm.MsgMigrateContract
	//MsgUpdateAdmin         = wasm.MsgUpdateAdmin
	//MsgClearAdmin          = wasm.MsgClearAdmin
	TIBCAcknowledgement        = tibcpacket.Acknowledgement
	NonFungibleTokenPacketData = tibctranfer.NonFungibleTokenPacketData
	MsgTIBCNftTransfer         = tibctranfer.MsgNftTransfer
	MsgTIBCUpdateClient        = tibcclient.MsgUpdateClient
	MsgTIBCRecvPacket          = tibcpacket.MsgRecvPacket
	MsgTIBCAcknowledgement     = tibcpacket.MsgAcknowledgement
	MsgCleanPacket             = tibcpacket.MsgCleanPacket
	MsgRecvCleanPacket         = tibcpacket.MsgRecvCleanPacket
)
