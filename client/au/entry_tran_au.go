package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// EntryTranAu is struct for EntryTranAu API.
// Entry for payment by au.
type EntryTranAu struct {
	client.BaseRequest `url:",squash"`

	OrderID string `url:"OrderID"`
	JobCd   string `url:"JobCd"`
	Amount  int64  `url:"Amount"`

	Tax         int64  `url:"Tax,omitempty"`
	PaymentType string `url:"PaymentType,omitempty"`
}

// Do executes EntryTranAu operation.
func (svc *EntryTranAu) Do(cli client.Client) (*EntryTranAuResponse, error) {
	const apiPath = "/payment/EntryTranAu.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	if svc.OrderID == "" {
		svc.OrderID = client.GetRandomOrderID()
	}

	result := &EntryTranAuResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// EntryTranAuResponse is struct for response of EntryTranAu API.
type EntryTranAuResponse struct {
	client.BaseResponse `url:",squash"`

	AccessID   string `url:"AccessID"`
	AccessPass string `url:"AccessPass"`
}

// IsSuccess checks the request is success or not
func (r *EntryTranAuResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.AccessID == "",
		r.AccessPass == "":
		return false
	}
	return true
}
