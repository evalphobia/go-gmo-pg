package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// AuAcceptStartResponse is struct for callback response of AuAcceptStartTranDate page.
type AuAcceptStartResponse struct {
	client.BaseResponse `url:",squash"`

	ShopID       string `url:"ShopID"`
	OrderID      string `url:"OrderID"`
	Status       string `url:"Status"`
	TranDate     string `url:"TranDate"`
	AuPayMethod  string `url:"AuPayMethod"`
	AuAcceptCode string `url:"AuAcceptCode"`
	CheckString  string `url:"CheckString"`
}

// IsSuccessRequest checks the request is success or not
func (r *AuAcceptStartResponse) IsSuccessRequest() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.Status == "":
		return false
	}
	return true
}

// IsSuccessStatus checks the request is success or not
func (r *AuAcceptStartResponse) IsSuccessStatus() bool {
	if !r.IsSuccessRequest() {
		return false
	}
	return r.Status == client.StatusRegister
}

// ValidateCheckString validates CheckString to avoid falsification.
func (r *AuAcceptStartResponse) ValidateCheckString(accessID, shopPass string) bool {
	return client.ValidateCheckString(r.CheckString, r.OrderID, accessID, r.ShopID, shopPass)
}
