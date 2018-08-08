package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// EntryTranAuAccept is struct for EntryTranAuAccept API.
// Entry for payment by au.
type EntryTranAuAccept struct {
	client.BaseRequest `url:",squash"`

	OrderID string `url:"OrderID"`
}

// Do executes EntryTranAuAccept operation.
func (svc *EntryTranAuAccept) Do(cli client.Client) (*EntryTranAuAcceptResponse, error) {
	const apiPath = "/payment/EntryTranAuAccept.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	if svc.OrderID == "" {
		svc.OrderID = client.GetRandomOrderID()
	}

	result := &EntryTranAuAcceptResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// EntryTranAuAcceptResponse is struct for response of EntryTranAuAccept API.
type EntryTranAuAcceptResponse struct {
	client.BaseResponse `url:",squash"`

	AccessID   string `url:"AccessID"`
	AccessPass string `url:"AccessPass"`
}

// IsSuccess checks the request is success or not
func (r *EntryTranAuAcceptResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.AccessID == "",
		r.AccessPass == "":
		return false
	}
	return true
}
