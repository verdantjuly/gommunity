package cache

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io/ioutil"
	"log"

	"github.com/andybalholm/brotli"
)

// Deflate 압축 함수
func CompressDeflate(input []byte) ([]byte, error) {
	var output bytes.Buffer
	writer, err := flate.NewWriter(&output, flate.BestCompression)
	if err != nil {
		return nil, err
	}
	defer writer.Close()

	_, err = writer.Write(input)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// Gzip 압축 함수
func CompressGzip(input []byte) ([]byte, error) {
	var output bytes.Buffer
	writer := gzip.NewWriter(&output)
	defer writer.Close()

	_, err := writer.Write(input)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// Brotli 압축 함수
func CompressBrotli(input []byte) ([]byte, error) {
	var output bytes.Buffer
	writer := brotli.NewWriter(&output)
	defer writer.Close()

	_, err := writer.Write(input)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// 압축 해제 함수 (Deflate)
func DecompressDeflate(input []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(input))
	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("Deflate 압축 해제 중 오류 발생:", err)
		return nil, err
	}

	return decompressed, nil
}

// 압축 해제 함수 (Gzip)
func DecompressGzip(input []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		log.Println("Gzip 압축 해제 중 오류 발생:", err)
		return nil, err
	}
	defer reader.Close()

	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("Gzip 압축 해제 중 오류 발생:", err)
		return nil, err
	}

	return decompressed, nil
}

// 압축 해제 함수 (Brotli)
func DecompressBrotli(input []byte) ([]byte, error) {
	reader := brotli.NewReader(bytes.NewReader(input))
	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("Brotli 압축 해제 중 오류 발생:", err)
		return nil, err
	}

	return decompressed, nil
}
