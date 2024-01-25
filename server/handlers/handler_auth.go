package handlers

import (
	"gommunity/server/api"
	"net/http"
)

func Authorization(w http.ResponseWriter, r *http.Request) api.Auth {
	headers := r.Header
	authorization := headers.Get("Authorization")
	VerifiedReuslt := api.Verify(authorization)
	return VerifiedReuslt
}
