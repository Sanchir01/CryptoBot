package myBinancePKG

import "github.com/adshao/go-binance/v2"

type BinanceStruct struct {
	binance *binance.Client
}

func NewBinanceClient(binance *binance.Client) *BinanceStruct {
	return &BinanceStruct{binance: binance}
}
