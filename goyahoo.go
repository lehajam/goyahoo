package goyahoo

/*
Implementation based on yoc (yahoo option chains) https://github.com/boyank/yoc
We don't panic and let clients decide what to do with errors
*/

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

var (
	url1 = "https://query{id}.finance.yahoo.com/v7/finance/options/{ticker}"
	url2 = "https://query{id}.finance.yahoo.com/v7/finance/options/{ticker}?date={expiry}"
)

func urlReplacer(ID, Symbol, Expiry string) *strings.Replacer {
	return strings.NewReplacer(
		"{id}", ID,
		"{ticker}", url.QueryEscape(Symbol),
		"{expiry}", Expiry)
}

type urlParams struct {
	QuerySymbol string
	QueryExpiry int64
}

func getQueryURL(QueryID string, Params urlParams) string {
	if Params.QueryExpiry != 0 {
		return urlReplacer(
			QueryID,
			Params.QuerySymbol,
			strconv.FormatInt(Params.QueryExpiry, 10)).Replace(url2)
	}

	return urlReplacer(
		QueryID,
		Params.QuerySymbol,
		"").Replace(url1)
}

// gojson -input AAPL.json -o go-optionchainroot.go -name optionChainRoot -pkg goyahoo
func getOptionChainRoot(Params urlParams) (*OptionChainRoot, error) {

	response, err := http.Get(getQueryURL("1", Params))
	if err != nil {
		response, err = http.Get(getQueryURL("2", Params))
		if err != nil {
			return nil, err
		}
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var root OptionChainRoot
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}
	if root.OptionChain.Error != nil {
		return nil, errors.New("")
	}

	return &root, nil
}

/*
GetExpiries returns all option expiries avaiable for a given symbol
*/
func GetExpiries(Symbol string) ([]int64, error) {
	root, err := getOptionChainRoot(urlParams{QuerySymbol: Symbol})
	if err != nil {
		return nil, err
	}

	return root.OptionChain.Result[0].ExpirationDates, nil
}

/*
GetExpiryChain returns the option chain from YAHOO finance for a given expiry and symbol
*/
func GetExpiryChain(Symbol string, Expiry int64) (*Chain, error) {
	root, err := getOptionChainRoot(urlParams{QuerySymbol: Symbol, QueryExpiry: Expiry})
	if err != nil {
		return nil, err
	}

	return &root.OptionChain.Result[0], nil
}

/*
GetOptionChain returns the option chain from YAHOO finance for the Symbol
It uses https://query1.finance.yahoo.com/v7/finance/options first
If it yields an error we try https://query2.finance.yahoo.com/v7/finance/options
If it also fails we return an error
We use QueryEscape from the package net/url to encode any special character in the ticker
*/
func GetOptionChain(Symbol string) ([]*Chain, []error, error) {
	expiries, err := GetExpiries(Symbol)
	if err != nil {
		return nil, nil, err
	}

	var wg sync.WaitGroup
	optionChain := make([]*Chain, len(expiries))
	chainErrors := make([]error, len(expiries))
	for index, expiry := range expiries {
		wg.Add(1)
		go func(i int, e int64) {
			chain, err := GetExpiryChain(Symbol, e)
			if err != nil {
				chainErrors[i] = err
			}
			optionChain[i] = chain
			wg.Done()
		}(index, expiry)
	}
	wg.Wait()
	return optionChain, chainErrors, nil
}
