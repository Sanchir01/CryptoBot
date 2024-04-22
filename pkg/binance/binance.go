package binance

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

type Binance struct {
	binance *binance.Client
}

func NewBinanceClient(binance *binance.Client) *Binance {
	return &Binance{binance: binance}
}

func (bin *Binance) initBinance() {

	prices, err := bin.binance.NewListPriceChangeStatsService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
}
