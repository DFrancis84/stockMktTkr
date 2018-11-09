package main

import (
	config "stockMkt/config"
	quote "stockMkt/quote"
)

var ()

func main() {
	key := config.GetKey()
	symbols := config.GetSymbols()
	quote.GetImmediateStockQuote(symbols, key)
}
