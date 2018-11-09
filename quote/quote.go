package quote

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	aaRoot = "https://www.alphavantage.co/query"
)

// API provides methods for interacting with AA
type API struct {
	symbols []string
	key     string
}

// New returns reference to AA API
func New(symbols []string, key string) *API {
	return &API{
		symbols: symbols,
		key:     key,
	}
}

func GetImmediateStockQuote(symbols []string, key string) {
	client := http.DefaultClient
	quote := API{
		symbols: symbols,
		key:     key,
	}

	for _, symbol := range symbols {
		getQuoteURI := fmt.Sprintf("%v?function=GLOBAL_QUOTE&symbol=%v&apikey=%v", aaRoot, symbol, quote.key)
		request, err := createDefaultRequest(http.MethodGet, getQuoteURI)
		if err != nil {
			log.Fatal(err)
		}

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		sB := string(body)
		fmt.Printf(sB)
		time.Sleep(12 * time.Second)
	}

}

func createDefaultRequest(method, uri string) (*http.Request, error) {
	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}
