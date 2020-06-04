package home

import (
	"net/http"

	"web20.tk/router/common"
)

func Home(w http.ResponseWriter, r *http.Request, context *map[string]interface{}) (string, *common.AppError) {
	return "frontend/dist/templates/home.tmpl", nil
}
