package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		shopID   string
		shopPass string
		hasError bool
	}{
		{"abcdefgh12345", "abcd1234", false},
		{"abcdefgh1234", "abcd1234", true},   // len(ShopID)=12
		{"abcdefgh12345", "abcd123", true},   // len(ShopPass)=7
		{"abcdefgh123456", "abcd1234", true}, // len(ShopID)=14
		{"abcdefgh12345", "abcd12345", true}, // len(ShopPass)=9
		{"", "", true},
		{"a", "b", true},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)
		conf, err := New(tt.shopID, tt.shopPass)
		switch {
		case tt.hasError:
			a.Error(err, target)
		default:
			a.NoError(err, target)
		}
		a.Equal(tt.shopID, conf.ShopID, target)
		a.Equal(tt.shopPass, conf.ShopPass, target)
	}
}

func TestValidate(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		shopID      string
		shopPass    string
		expectedErr string
	}{
		{"abcdefgh12345", "abcd1234", ""}, // no error
		{"abcdefgh1234", "abcd1234", "[ShopID] must be 13 length char"},
		{"abcdefgh12345", "abcd123", "[ShopPass] must be 8 length char"},
		{"abcdefgh123456", "abcd1234", "[ShopID] must be 13 length char"},
		{"abcdefgh12345", "abcd12345", "[ShopPass] must be 8 length char"},
		{"", "", "[ShopID] is mandatory | [ShopPass] is mandatory"},
		{"", "abcd1234", "[ShopID] is mandatory"},
		{"abcdefgh12345", "", "[ShopPass] is mandatory"},
		{"a", "b", "[ShopID] must be 13 length char | [ShopPass] must be 8 length char"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)
		conf := Config{
			ShopID:   tt.shopID,
			ShopPass: tt.shopPass,
		}
		err := conf.Validate()
		switch {
		case tt.expectedErr == "":
			a.NoError(err, target)
		default:
			a.Equal(tt.expectedErr, err.Error(), target)
		}
	}
}
