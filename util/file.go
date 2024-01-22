package util

import (
	"errors"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func ContentType(path string) string {
	return mime.TypeByExtension(filepath.Ext(path))
}

func ContentAndType(path string) ([]byte, string) {
	content, err := os.ReadFile(path)

	// 파일이 존재하지 않는 경우를 제외합니다.
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Println("파일을 불러오는 과정에서 예외가 발생하였습니다.", err.Error())
	}

	return content, ContentType(path)
}

func JoinPath(elements ...string) string {
	return filepath.Join(elements...)
}

func GetFileExtension(filePath string) string {
	// 파일 경로에서 마지막 "." 이후의 문자열을 추출하여 파일 확장자로 사용
	dotIndex := strings.LastIndex(filePath, ".")
	if dotIndex != -1 {
		return filePath[dotIndex+1:]
	}
	return ""
}

func GetBaseDir() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(executable), nil
}
