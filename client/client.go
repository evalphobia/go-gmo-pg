package client

import (
	"fmt"

	"github.com/evalphobia/go-gmo-pg/config"
	"github.com/evalphobia/go-gmo-pg/request"
)

const (
	endpointBase = "https://pt01.mul-pay.jp"
)

// Client is base struct for GMO Payment Gateway API.
type Client struct {
	Config *config.Config `url:"-"`
	Option request.Option `url:"-"`
}

// New creates Client with given config.
func New(conf *config.Config) Client {
	return Client{
		Config: conf,
	}
}

// Call sends HTTP request to GMO Payment Gateway API.
func (c Client) Call(path string, param interface{}, result interface{}) error {
	p := parameter{
		Common: c,
		Extra:  param,
	}
	return request.CallPOST(fmt.Sprintf("%s%s", endpointBase, path), p, c.Option, result)
}

type parameter struct {
	Common Client      `url:",squash"`
	Extra  interface{} `url:",squash"`
}
