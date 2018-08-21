package config

import (
	"errors"
	"strings"
)

const (
	modeSandbox = iota + 1
	modeProduction
)

// Config is struct for base setting for GMO Payment Gateway API.
type Config struct {
	Mode     `url:"-"`
	Version  string `url:"Version,omitempty"`
	ShopID   string `url:"ShopID"`
	ShopPass string `url:"ShopPass"`
}

// New returns initialized *Config
func New(id, pass string) (*Config, error) {
	c := &Config{
		Mode:     modeSandbox,
		ShopID:   id,
		ShopPass: pass,
	}
	return c, c.Validate()
}

// Validate validates parameters.
func (c Config) Validate() error {
	var errList []string
	switch {
	case c.ShopID == "":
		errList = append(errList, "[ShopID] is mandatory")
	case len(c.ShopID) != 13:
		errList = append(errList, "[ShopID] must be 13 length char")
	}

	switch {
	case c.ShopPass == "":
		errList = append(errList, "[ShopPass] is mandatory")
	case len(c.ShopPass) != 8:
		errList = append(errList, "[ShopPass] must be 8 length char")
	}

	if len(errList) == 0 {
		return nil
	}
	return errors.New(strings.Join(errList, " | "))
}

// SetAsProduction set mode to production.
func (c *Config) SetAsProduction() {
	c.Mode = modeProduction
}

// Mode stands for development mode or live mode.
type Mode int

// IsProduction checks the mode is production(=live) or not.
func (m Mode) IsProduction() bool {
	return m == modeProduction
}
