package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// AuAcceptUserEnd is struct for AuAcceptUserEnd API.
// Entry for payment by au.
type AuAcceptUserEnd struct {
	client.BaseRequest `url:",squash"`

	AccessID     string `url:"AccessID"`
	AccessPass   string `url:"AccessPass"`
	OrderID      string `url:"OrderID"`
	AuAcceptCode string `url:"AuAcceptCode"`
}

// Do executes AuAcceptUserEnd operation.
func (svc *AuAcceptUserEnd) Do(cli client.Client) (*AuAcceptUserEndResponse, error) {
	const apiPath = "/payment/AuAcceptUserEnd.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	if svc.OrderID == "" {
		svc.OrderID = client.GetRandomOrderID()
	}

	result := &AuAcceptUserEndResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// AuAcceptUserEndResponse is struct for response of AuAcceptUserEnd API.
type AuAcceptUserEndResponse struct {
	client.BaseResponse `url:",squash"`

	OrderID string `url:"OrderID"`
	Status  string `url:"Status"`
}

// IsSuccessRequest checks the request is success or not
func (r *AuAcceptUserEndResponse) IsSuccessRequest() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.Status == "":
		return false
	}
	return true
}

// IsSuccessStatus checks the request is success or not
func (r *AuAcceptUserEndResponse) IsSuccessStatus() bool {
	if !r.IsSuccessRequest() {
		return false
	}
	return r.Status == client.StatusEnd
}
