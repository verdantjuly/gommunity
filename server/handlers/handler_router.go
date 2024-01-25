package handlers

import (
	"gommunity/server/cache"
	"gommunity/server/utils"
	"net/http"
	"path/filepath"
)

type HTTPHandler struct {
	http.Handler
}

func (HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGetRequest(w, r)
	}
	if r.Method == http.MethodPost {
		handlePostRequest(w, r)
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	res := cache.Resource{}
	ext := filepath.Ext(r.URL.Path)

	if ext == "" && r.URL.Path == "/" {
		res.Path = utils.JoinPath("../client/templates", "index.html")

	} else if ext == "" && r.URL.Path != "/" {
		url := r.URL.Path + ".html"
		res.Path = utils.JoinPath("../client/templates", url)
	} else {
		res.Path = utils.JoinPath("../client/templates", r.URL.Path)

	}

	// 캐시된 데이터 로드
	data := res.Load()

	// 캐시된 데이터가 없으면 파일 읽어와서 캐싱
	if data == nil {
		content, contentType := utils.ContentAndType(res.Path)
		if content == nil {
			http.NotFound(w, r)
			return
		}

		data = &cache.ResourceData{
			Parent:      &res,
			Content:     content,
			ContentType: contentType,
		}

		// 캐시에 데이터 추가
		res.Cache(data)
	}

	// 캐시된 데이터 전송
	data.WriteWith(w)
}
func handlePostRequest(w http.ResponseWriter, r *http.Request) {

	var data cache.ResourceData

	switch r.URL.Path {
	case "/signup":
		data = SignUpPostRequest(w, r)

	}
	data.WriteWith(w)
}
