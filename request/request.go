package request

import (
	"strings"

	"github.com/evalphobia/httpwrapper/request"
	"github.com/mitchellh/mapstructure"
)

// CallPOST sends POST request to `url` with `params` and set reqponse to `result`
func CallPOST(url string, params interface{}, opt Option, result interface{}) (err error) {
	resp, err := request.POST(url, request.Option{
		Payload:     params,
		PayloadType: request.PayloadTypeFORM,
		Retry:       opt.Retry,
		Debug:       opt.Debug,
		UserAgent:   opt.getUserAgent(),
		Timeout:     opt.getTimeout(),
	})
	if err != nil {
		return err
	}

	mapData := parseToMap(resp.String())
	switch v := result.(type) {
	case unmarshaler:
		err = assignFromMap(mapData, result)
		if err != nil {
			return err
		}
		v.Unmarshal(mapData)
		return nil
	default:
		return assignFromMap(mapData, result)
	}
}

// parseToMap converts response string data to map.
func parseToMap(str string) map[string]interface{} {
	m := make(map[string]interface{})
	values := strings.Split(str, "&")
	for _, value := range values {
		v := strings.Split(value, "=")
		if len(v) != 2 {
			continue
		}
		m[v[0]] = v[1]
	}
	return m
}

// assignFromMap set data from map to struct.
func assignFromMap(mapData interface{}, result interface{}) error {
	const tagName = "url"
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   result,
		TagName:  tagName,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(mapData)
}

type unmarshaler interface {
	Unmarshal(mapData map[string]interface{}) error
}
