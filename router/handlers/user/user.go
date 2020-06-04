package user

import (
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return "frontend/dist/templates/user.tmpl", nil
}
