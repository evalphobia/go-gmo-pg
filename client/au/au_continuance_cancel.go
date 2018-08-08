package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// AuContinuanceCancel is struct for AuContinuanceCancel API.
// Cancel subscription payment by au.
type AuContinuanceCancel struct {
	client.BaseRequest `url:",squash"`

	// required
	OrderID    string `url:"OrderID"`
	AccessID   string `url:"AccessID"`
	AccessPass string `url:"AccessPass"`
}

// Do executes AuContinuanceCancel operation.
func (svc *AuContinuanceCancel) Do(cli client.Client) (*AuContinuanceCancelResponse, error) {
	const apiPath = "/payment/AuContinuanceCancel.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	result := &AuContinuanceCancelResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// AuContinuanceCancelResponse is struct for response of AuContinuanceCancel API.
type AuContinuanceCancelResponse struct {
	client.BaseResponse `url:",squash"`

	OrderID string `url:"OrderID"`
	Status  string `url:"Status"`
}

// IsSuccess checks the request is success or not
func (r *AuContinuanceCancelResponse) IsSuccess() bool {
	if !r.BaseResponse.IsSuccess() {
		return false
	}
	return r.Status == client.StatusCancel
}
