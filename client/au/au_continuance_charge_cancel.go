package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// AuContinuanceChargeCancel is struct for AuContinuanceChargeCancel API.
// Cancel subscription payment by au.
type AuContinuanceChargeCancel struct {
	client.BaseRequest `url:",squash"`

	// required
	OrderID          string `url:"OrderID"`
	AccessID         string `url:"AccessID"`
	AccessPass       string `url:"AccessPass"`
	CancelAmount     int64  `url:"CancelAmount"`
	CancelTax        int64  `url:"CancelTax"`
	ContinuanceMonth string `url:"ContinuanceMonth"`
}

// Do executes AuContinuanceChargeCancel operation.
func (svc *AuContinuanceChargeCancel) Do(cli client.Client) (*AuContinuanceChargeCancelResponse, error) {
	const apiPath = "/payment/AuContinuanceChargeCancel.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	result := &AuContinuanceChargeCancelResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// AuContinuanceChargeCancelResponse is struct for response of AuContinuanceChargeCancel API.
type AuContinuanceChargeCancelResponse struct {
	client.BaseResponse `url:",squash"`

	OrderID          string `url:"OrderID"`
	Status           string `url:"Status"`
	ContinuanceMonth string `url:"ContinuanceMonth"`
	Amount           int64  `url:"Amount"`
	Tax              int64  `url:"Tax"`
	CancelAmount     int64  `url:"CancelAmount"`
	CancelTax        int64  `url:"CancelTax"`
}

// IsSuccess checks the request is success or not
func (r *AuContinuanceChargeCancelResponse) IsSuccess() bool {
	if !r.BaseResponse.IsSuccess() {
		return false
	}
	switch r.Status {
	case client.StatusCancel,
		client.StatusReturn:
		return true
	}
	return false
}
