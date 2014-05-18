package server

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func ImageHandler() *Handler {
	handler := &Handler{
		Serve: imageServe,
	}
	return handler
}

func imageServe(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		imageCreateServe(writer, request)
	}
}

func imageCreateServe(writer http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("image")

	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	err = ioutil.WriteFile("/tmp/"+filename(handler.Filename), data, 0777)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Write([]byte(filename(handler.Filename)))
}

func filename(filename string) string {
	hash := md5.New()
	hash.Write([]byte(filename))
	return hex.EncodeToString(hash.Sum(nil)) + filepath.Ext(filename)
}
