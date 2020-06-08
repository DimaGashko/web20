package md

import (
	"encoding/json"
	"net/http"

	"web20.tk/entries"
	"web20.tk/router/handlers/common"
)

func Md2Html(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	resp := make(map[string]interface{})
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return resp, common.WrapError(err, http.StatusBadRequest)
	}

	md, ok := body["md"]
	if !ok {
		return resp, common.New("'md' field is empty", http.StatusBadRequest)
	}

	resp["html"] = entries.Md(md.(string)).Html()
	return resp, nil
}
