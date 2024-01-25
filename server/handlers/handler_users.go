package handlers

import (
	"gommunity/server/cache"
	"net/http"
)

func SignUpPostRequest(w http.ResponseWriter, r *http.Request) cache.ResourceData {

	return cache.ResourceData{
		ContentType: "application/json",
	}

}
