package main

import (
	"stockMkt/config"
	"stockMkt/quote"
)

var ()

func main() {
	key := config.GetKey()
	symbols := config.GetSymbols()
	quote.GetImmediateStockQuote(symbols, key)
}
