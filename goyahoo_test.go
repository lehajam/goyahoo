package goyahoo

import (
	"testing"
)

/*
Test the query URL generation
*/
func Test_getQueryURL(t *testing.T) {
	type args struct {
		QueryID string
		Params  urlParams
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"AAPL(Apple) query1 URL",
			args{QueryID: "1", Params: urlParams{QuerySymbol: "AAPL"}},
			"https://query1.finance.yahoo.com/v7/finance/options/AAPL",
		},
		{
			"AAPL(Apple) query2 URL",
			args{QueryID: "2", Params: urlParams{QuerySymbol: "AAPL"}},
			"https://query2.finance.yahoo.com/v7/finance/options/AAPL",
		},
		{
			"AAPL(Apple) with expiry(epoch time) 1513296000 query1 URL",
			args{QueryID: "1", Params: urlParams{QuerySymbol: "AAPL", QueryExpiry: 1513296000}},
			"https://query1.finance.yahoo.com/v7/finance/options/AAPL?date=1513296000",
		},
		{
			"AAPL(Apple) with expiry(epoch time) 1513296000 query2 URL",
			args{QueryID: "2", Params: urlParams{QuerySymbol: "AAPL", QueryExpiry: 1513296000}},
			"https://query2.finance.yahoo.com/v7/finance/options/AAPL?date=1513296000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getQueryURL(tt.args.QueryID, tt.args.Params); got != tt.want {
				t.Errorf("getQueryURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOptionChainRoot(t *testing.T) {
}

func TestGetExpiries(t *testing.T) {
	if _, err := GetExpiries("AAPL"); err != nil {
		t.Error(err)
	}
}

func TestGetExpiryChain(t *testing.T) {
	expiries, err := GetExpiries("AAPL")
	if err != nil {
		t.Error(err)
	}

	for _, expiry := range expiries {
		if _, err := GetExpiryChain("AAPL", expiry); err != nil {
			t.Error(err)
		}
	}
}

func TestGetOptionChain(t *testing.T) {
	_, chainErrors, err := GetOptionChain("AAPL")
	if err != nil {
		t.Error(err)
	}
	for _, err := range chainErrors {
		if err != nil {
			t.Error(err)
		}
	}
}
