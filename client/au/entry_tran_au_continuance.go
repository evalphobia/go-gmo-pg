package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// EntryTranAuContinuance is struct for EntryTranAuContinuance API.
// Entry for subscription payment by au.
type EntryTranAuContinuance struct {
	client.BaseRequest `url:",squash"`

	OrderID     string `url:"OrderID"`
	FirstAmount int64  `url:"FirstAmount"`
	Amount      int64  `url:"Amount"`
	FirstTax    int64  `url:"FirstTax,omitempty"`
	Tax         int64  `url:"Tax,omitempty"`
}

// Do executes EntryTranAuContinuance operation.
func (svc *EntryTranAuContinuance) Do(cli client.Client) (*EntryTranAuContinuanceResponse, error) {
	const apiPath = "/payment/EntryTranAuContinuance.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	if svc.OrderID == "" {
		svc.OrderID = client.GetRandomOrderID()
	}

	result := &EntryTranAuContinuanceResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// EntryTranAuContinuanceResponse is struct for response of EntryTranAuContinuance API.
type EntryTranAuContinuanceResponse struct {
	client.BaseResponse `url:",squash"`

	AccessID   string `url:"AccessID"`
	AccessPass string `url:"AccessPass"`
}

// IsSuccess checks the request is success or not
func (r *EntryTranAuContinuanceResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.AccessID == "",
		r.AccessPass == "":
		return false
	}
	return true
}
