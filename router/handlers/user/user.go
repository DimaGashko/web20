package user

import (
	"net/http"

	"web20.tk/router/common"
)

func User(w http.ResponseWriter, r *http.Request, context *map[string]interface{}) (string, *common.AppError) {
	return "frontend/dist/templates/user.tmpl", nil
}
