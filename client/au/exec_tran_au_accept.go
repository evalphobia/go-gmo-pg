package au

import (
	"time"

	"github.com/evalphobia/go-gmo-pg/client"
)

// ExecTranAuAccept is struct for ExecTranAuAccept API.
// Execution for payment by au.
type ExecTranAuAccept struct {
	client.BaseRequest `url:",squash"`

	// required
	OrderID     string `url:"OrderID"`
	AccessID    string `url:"AccessID"`
	AccessPass  string `url:"AccessPass"`
	Commodity   string `url:"Commodity"`
	RetURL      string `url:"RetURL"`
	ServiceName string `url:"ServiceName"`
	ServiceTel  string `url:"ServiceTel"`

	// required when to set menber id
	MemberID     string `url:"MemberID,omitempty"`
	SiteID       string `url:"SiteID,omitempty"`
	SitePass     string `url:"SitePass,omitempty"`
	CreateMember string `url:"CreateMember,omitempty"`

	// optional
	MemberName     string `url:"MemberName,omitempty"`
	PaymentTermSec int    `url:"PaymentTermSec,omitempty"`
	ClientField1   string `url:"ClientField1,omitempty"`
	ClientField2   string `url:"ClientField2,omitempty"`
	ClientField3   string `url:"ClientField3,omitempty"`
}

// Do executes ExecTranAuAccept operation.
func (svc *ExecTranAuAccept) Do(cli client.Client) (*ExecTranAuAcceptResponse, error) {
	const apiPath = "/payment/ExecTranAuAccept.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	result := &ExecTranAuAcceptResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// DoWithSjis executes ExecTranAuAccept operation with Shift_JIS encoding.
func (svc *ExecTranAuAccept) DoWithSjis(cli client.Client) (*ExecTranAuAcceptResponse, error) {
	sjisData, err := svc.ToSjis()
	if err != nil {
		return nil, err
	}
	return sjisData.Do(cli)
}

// ToSjis uses parameters as Shift_JIS.
func (svc ExecTranAuAccept) ToSjis() (ExecTranAuAccept, error) {
	var err error
	sjisData := svc

	sjisData.ServiceName, err = client.ConvertUtf8ToSjis(client.ConvertToFullWidth(svc.ServiceName))
	if err != nil {
		return sjisData, err
	}
	sjisData.Commodity, err = client.ConvertUtf8ToSjis(client.ConvertToFullWidth(svc.Commodity))
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

// ExecTranAuAcceptResponse is struct for response of ExecTranAuAccept API.
type ExecTranAuAcceptResponse struct {
	client.BaseResponse `url:",squash"`

	AccessID       string `url:"AccessID"`
	Token          string `url:"Token"`
	StartURL       string `url:"StartURL"`
	StartLimitDate string `url:"StartLimitDate"`
}

// GetStartLimitDate returns StartLimitDate in time.Time type.
func (r ExecTranAuAcceptResponse) GetStartLimitDate() time.Time {
	dt, _ := time.Parse("20060102150405", r.StartLimitDate)
	return dt
}

// IsSuccess checks the request is success or not
func (r *ExecTranAuAcceptResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.AccessID == "",
		r.Token == "":
		return false
	}
	return true
}
