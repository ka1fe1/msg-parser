module github.com/kaifei-bianjie/msg-parser

go 1.14

require (
	github.com/bianjieai/iritamod v0.0.0-20211201070223-cce51303d50f
	github.com/cosmos/cosmos-sdk v0.40.0-rc3
	github.com/irisnet/irismod v1.1.1-0.20201211020601-9c939d7f8ccc
	github.com/jolestar/go-commons-pool v2.0.0+incompatible // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc6
	gitlab.bianjie.ai/cschain/cschain v1.1.1-0.20211201072800-2b6efc61bc2c
	go.uber.org/zap v1.15.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.28.2-0.20210112055458-b53a7d5a7c9c
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104.0.20210112015006-57e95aa6402f
)
