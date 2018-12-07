package quote

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	avRoot = "https://www.alphavantage.co/query"
)

// API provides methods for interacting with Alpha Vantage
type API struct {
	symbols []string
	key     string
}

// New returns reference to Alpha Vantage API
func New(symbols []string, key string) *API {
	return &API{
		symbols: symbols,
		key:     key,
	}
}

// GetImmediateStockQuote sends request to Alpha Vantage and returns stock quotes based on stored symbols in config file.
func GetImmediateStockQuote(symbols []string, key string) {
	client := http.DefaultClient
	quote := API{
		symbols: symbols,
		key:     key,
	}

	for _, symbol := range symbols {
		getQuoteURI := fmt.Sprintf("%v?function=GLOBAL_QUOTE&symbol=%v&apikey=%v", avRoot, symbol, quote.key)
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

		// Alpha Vantage allows 5 request per minute, this timer will limit to 5 requests per minute if over 5 symbols
		if len(symbols) > 5 {
			time.Sleep(12 * time.Second)
		}
	}
}

func createDefaultRequest(method, uri string) (*http.Request, error) {
	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}
