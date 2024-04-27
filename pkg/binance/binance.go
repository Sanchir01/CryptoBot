package myBinancePKG

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

type BinanceStruct struct {
	binance *binance.Client
}

func NewBinanceClient(binance *binance.Client) *BinanceStruct {
	return &BinanceStruct{binance: binance}
}

func (bina *BinanceStruct) GetAllPricesCryptocurrency() ([]*binance.SymbolPrice, error) {
	prices, err := bina.binance.NewListPricesService().Do(context.Background())
	fmt.Printf("this is binance bitcoin value %v", bina.binance.NewListPricesService().Symbol("BTCUSDT"))
	fmt.Printf("this is binance bitcoin value %v", bina.binance.NewListPricesService().Symbol("ETHUSDT"))

	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (bina *BinanceStruct) GetOneCryptocurrency(name string) (*binance.ListPricesService, error) {
	price := bina.binance.NewListPricesService().Symbol(name)

	return price, nil
}
