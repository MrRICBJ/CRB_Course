package client

import (
	"net/http"
	"time"
)

// FetchFunction is a function that mimics http.Get() method
type FetchFunction func(url string) (resp *http.Response, err error)

// Client is a currency rates service client... what else?
type Client interface {
	GetRate(time.Time) (Result, error)
	SetFetchFunction(FetchFunction)
}

type client struct {
	fetch FetchFunction
}

func (s client) GetRate(t time.Time) (Result, error) {
	rate, err := getCurrencyRate(t, s.fetch)
	if err != nil {
		return Result{}, err
	}
	return rate, nil
}

func (s client) SetFetchFunction(f FetchFunction) {
	s.fetch = f
}

// NewClient creates a new rates service instance
func NewClient() Client {
	return client{http.Get}
}
