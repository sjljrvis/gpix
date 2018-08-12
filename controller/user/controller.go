package controller

import (
	helper "github.com/sjljrvis/gpix/helper"
	"net/http"
)

// HomePage controller
func HomePage(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, 200, map[string]string{"user": "sejal"})
}
