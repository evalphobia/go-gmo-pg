package au

import (
	"github.com/evalphobia/go-gmo-pg/client"
)

// ExecTranAu is struct for ExecTranAu API.
// Execution for payment by au.
type ExecTranAu struct {
	client.BaseRequest `url:",squash"`

	// required
	OrderID     string `url:"OrderID"`
	AccessID    string `url:"AccessID"`
	AccessPass  string `url:"AccessPass"`
	Commodity   string `url:"Commodity"`
	ServiceName string `url:"ServiceName"`
	ServiceTel  string `url:"ServiceTel"`

	AuAcceptCode string `url:"AuAcceptCode"`

	// optional
	ClientField1 string `url:"ClientField1,omitempty"`
	ClientField2 string `url:"ClientField2,omitempty"`
	ClientField3 string `url:"ClientField3,omitempty"`
}

// Do executes ExecTranAu operation.
func (svc *ExecTranAu) Do(cli client.Client) (*ExecTranAuResponse, error) {
	const apiPath = "/payment/ExecTranAu.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	result := &ExecTranAuResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// DoWithSjis executes ExecTranAu operation with Shift_JIS encoding.
func (svc *ExecTranAu) DoWithSjis(cli client.Client) (*ExecTranAuResponse, error) {
	sjisData, err := svc.ToSjis()
	if err != nil {
		return nil, err
	}
	return sjisData.Do(cli)
}

// ToSjis uses parameters as Shift_JIS.
func (svc ExecTranAu) ToSjis() (ExecTranAu, error) {
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

// ExecTranAuResponse is struct for response of ExecTranAu API.
type ExecTranAuResponse struct {
	client.BaseResponse `url:",squash"`

	OrderID     string `url:"OrderID"`
	Status      string `url:"Status"`
	TranDate    string `url:"TranDate"`
	PayInfoNo   string `url:"PayInfoNo"`
	PayMethod   string `url:"PayMethod"`
	CheckString string `url:"CheckString"`
}

// IsSuccess checks the request is success or not
func (r *ExecTranAuResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.Status == "":
		return false
	}
	return true
}

// ValidateCheckString validates CheckString to avoid falsification.
func (r *ExecTranAuResponse) ValidateCheckString(accessID, shopID, shopPass string) bool {
	return client.ValidateCheckString(r.CheckString, r.OrderID, accessID, shopID, shopPass)
}
