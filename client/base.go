package client

// BaseRequest is base struct for api request.
type BaseRequest struct {
	Version  string `url:"Version,omitempty" json:"-"`
	ShopID   string `url:"ShopID,omitempty" json:"-"`
	ShopPass string `url:"ShopPass,omitempty" json:"-"`
}

// BaseResponse is base struct for api response.
type BaseResponse struct {
	ErrCode string `url:"ErrCode"`
	ErrInfo string `url:"ErrInfo"`
}

// IsSuccess checks the request is success or not.
func (r BaseResponse) IsSuccess() bool {
	return r.ErrCode == "" && r.ErrInfo == ""
}

// Error returns error message
func (r BaseResponse) Error() string {
	return r.ErrInfo
}

// Response is interface of response from each API.
type Response interface {
	IsSuccess() bool
	Error() string
}
