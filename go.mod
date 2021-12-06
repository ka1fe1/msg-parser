module github.com/kaifei-bianjie/msg-parser

go 1.15

require (
	github.com/bianjieai/iritamod v1.0.1-0.20211026024119-806ffe494860
	github.com/bianjieai/tibc-go v0.2.0-alpha
	github.com/cosmos/cosmos-sdk v0.44.2
	github.com/irisnet/irismod v1.4.1-0.20211021075334-969a56e99ce9
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.13
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
)
