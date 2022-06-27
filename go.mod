module github.com/kaifei-bianjie/msg-parser

go 1.14

require (
	github.com/bianjieai/iritamod v1.2.1-0.20220222035322-99168809cf24
	github.com/cosmos/cosmos-sdk v0.44.4
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/irismod v1.5.2
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.35.0
	gitlab.bianjie.ai/cschain/cschain v1.1.1-0.20220701022148-93d0a666edb5
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104.0.20210112015006-57e95aa6402f
)
