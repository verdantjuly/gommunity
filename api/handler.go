package api

import (
	"gommunity/server/cache"
	"gommunity/server/util"
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
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	res := cache.Resource{}
	ext := filepath.Ext(r.URL.Path)

	if ext == "" {
		res.Path = util.JoinPath("../templates", "index.html")
	} else {
		res.Path = util.JoinPath("../templates", r.URL.Path)
	}

	// 캐시된 데이터 로드
	data := res.Load()

	// 캐시된 데이터가 없으면 파일 읽어와서 캐싱
	if data == nil {
		content, contentType := util.ContentAndType(res.Path)
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
