package cache

import (
	"gommunity/server/util"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Resource struct {
	Type string
	Path string
}

type ResourceData struct {
	Parent      *Resource
	Content     []byte
	ContentType string
}

var (
	cachedMutex = sync.RWMutex{}
	cached      = map[string]map[string]*ResourceData{}
)

func (r *Resource) Load() *ResourceData {
	cachedMutex.RLock()
	defer cachedMutex.RUnlock()

	return cached[r.Type][r.Path]
}

func (r *Resource) Cache(data *ResourceData) {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	if _, ok := cached[r.Type]; !ok {
		cached[r.Type] = make(map[string]*ResourceData)
	}
	cached[r.Type][r.Path] = data
}

func (r *Resource) WriteWith(w http.ResponseWriter) {
	if r.Path == "" {
		log.Println("주어진 캐싱 대상의 리소스 파일 경로가 정의되지 않았습니다")
		return
	}
	content, contentType := util.ContentAndType(r.Path)

	if content == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := ResourceData{
		Parent:      r,
		Content:     content,
		ContentType: contentType,
	}

	// 파일 유형에 따라 압축 방식 선택
	compressedData, err := compressBasedOnFileType(content, r.Path)
	if err != nil {
		log.Println("압축 중 오류 발생:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data.Content = compressedData

	// ResourceData의 WriteWith 메서드 호출
	data.WriteWith(w)
}

func (data *ResourceData) CompressType() string {
	return data.Parent.Type
}

func (data *ResourceData) WriteWith(w http.ResponseWriter) {
	w.Header().Set("Content-Encoding", data.CompressType())
	w.Header().Set("Content-Type", data.ContentType)
	w.Write(data.Content)
}

func compressBasedOnFileType(data []byte, filePath string) ([]byte, error) {
	// 파일 확장자 추출
	fileExt := getFileExtension(filePath)

	// 파일 확장자에 따라 압축 방식 선택
	switch fileExt {
	case "br":
		return CompressBrotli(data)
	case "gz":
		return CompressGzip(data)
	case "deflate":
		return CompressDeflate(data)
	default:
		// 기본적으로 Brotli 압축 사용
		return CompressBrotli(data)
	}
}

func getFileExtension(filePath string) string {
	// 파일 경로에서 마지막 "." 이후의 문자열을 추출하여 파일 확장자로 사용
	dotIndex := strings.LastIndex(filePath, ".")
	if dotIndex != -1 {
		return filePath[dotIndex+1:]
	}
	return ""
}
