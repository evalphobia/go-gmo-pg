package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// ExecTranAuContinuance is struct for ExecTranAuContinuance API.
// Execution for subscription payment by au.
type ExecTranAuContinuance struct {
	client.BaseRequest `url:",squash"`

	// required
	OrderID          string `url:"OrderID"`
	AccessID         string `url:"AccessID"`
	AccessPass       string `url:"AccessPass"`
	Commodity        string `url:"Commodity"`
	FirstAccountDate string `url:"FirstAccountDate"`
	RetURL           string `url:"RetURL"`
	ServiceName      string `url:"ServiceName"`
	ServiceTel       string `url:"ServiceTel"`

	IsAccountTimingEndOfMonth bool   `url:"-"`
	AccountTimingKbn          string `url:"AccountTimingKbn"`
	AccountTiming             string `url:"CreateMember,omitempty"`

	// required when to set menber id
	MemberID     string `url:"MemberID,omitempty"`
	SiteID       string `url:"SiteID,omitempty"`
	SitePass     string `url:"SitePass,omitempty"`
	CreateMember string `url:"CreateMember,omitempty"`

	// optional
	PaymentTermSec int    `url:"PaymentTermSec,omitempty"`
	ClientField1   string `url:"ClientField1,omitempty"`
	ClientField2   string `url:"ClientField2,omitempty"`
	ClientField3   string `url:"ClientField3,omitempty"`
}

// Do executes ExecTranAuContinuance operation.
func (svc *ExecTranAuContinuance) Do(cli client.Client) (*ExecTranAuContinuanceResponse, error) {
	const apiPath = "/payment/ExecTranAuContinuance.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	if svc.AccountTimingKbn == "" {
		switch {
		case svc.IsAccountTimingEndOfMonth:
			svc.AccountTimingKbn = client.AuAccountTimingEndOfMonth
		default:
			svc.AccountTimingKbn = client.AuAccountTimingSelect
		}
	}

	result := &ExecTranAuContinuanceResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// DoWithSjis executes ExecTranAuContinuance operation with Shift_JIS encoding.
func (svc *ExecTranAuContinuance) DoWithSjis(cli client.Client) (*ExecTranAuContinuanceResponse, error) {
	sjisData, err := svc.ToSjis()
	if err != nil {
		return nil, err
	}
	return sjisData.Do(cli)
}

// ToSjis uses parameters as Shift_JIS.
func (svc ExecTranAuContinuance) ToSjis() (ExecTranAuContinuance, error) {
	var err error
	sjisData := svc

	sjisData.ServiceName, err = client.ConvertUtf8ToSjis(client.ConvertToFullWidth(svc.ServiceName))
	if err != nil {
		return sjisData, err
	}

	commodity := trimCommodityUnderLimit(svc.Commodity)
	sjisData.Commodity, err = client.ConvertUtf8ToSjis(client.ConvertToFullWidth(commodity))
	if err != nil {
		return sjisData, err
	}
	sjisData.ClientField1, err = client.ConvertUtf8ToSjis(svc.ClientField1)
	if err != nil {
		return sjisData, err
	}
	sjisData.ClientField2, err = client.ConvertUtf8ToSjis(svc.ClientField2)
	if err != nil {
		return sjisData, err
	}
	sjisData.ClientField3, err = client.ConvertUtf8ToSjis(svc.ClientField3)
	return sjisData, err
}

// ExecTranAuContinuanceResponse is struct for response of ExecTranAuContinuance API.
type ExecTranAuContinuanceResponse struct {
	client.BaseResponse `url:",squash"`

	AccessID       string `url:"AccessID"`
	Token          string `url:"Token"`
	StartURL       string `url:"StartURL"`
	StartLimitDate string `url:"StartLimitDate"`
}

// IsSuccess checks the request is success or not
func (r *ExecTranAuContinuanceResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.AccessID == "",
		r.Token == "":
		return false
	}
	return true
}
