package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// AuContinuanceStartResponse is struct for callback response of AuContinuanceStart page.
type AuContinuanceStartResponse struct {
	client.BaseResponse `url:",squash"`

	ShopID             string `url:"ShopID"`
	OrderID            string `url:"OrderID"`
	Status             string `url:"Status"`
	TranDate           string `url:"TranDate"`
	AuContinuAccountId string `url:"AuContinuAccountId"`
	AuPayMethod        string `url:"AuPayMethod"`
}

// IsSuccessRequest checks the request is success or not
func (r *AuContinuanceStartResponse) IsSuccessRequest() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.Status == "":
		return false
	}
	return true
}

// IsSuccessStatus checks the request is success or not
func (r *AuContinuanceStartResponse) IsSuccessStatus() bool {
	if !r.IsSuccessRequest() {
		return false
	}
	return r.Status == client.StatusRegister
}
