package json

import "encoding/json"

func JsonResp(code int, data interface{}) string {
	resp := make(map[string]interface{})
	resp["code"] = code
	resp["data"] = data

	js, err := json.Marshal(resp)
	if err != nil {
		return err.Error()
	}
	return string(js)
}
