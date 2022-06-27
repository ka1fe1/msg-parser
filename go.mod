module github.com/kaifei-bianjie/msg-parser

go 1.15

require (
	github.com/bianjieai/iritamod v1.2.1-0.20220222035322-99168809cf24
	github.com/bianjieai/tibc-go v0.2.0-alpha
	github.com/cosmos/cosmos-sdk v0.44.4
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/irismod v1.5.2-0.20220222061735-b318859ba444
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.8.1
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.8.2-0.20220211020007-9ec25dde74d4
)
