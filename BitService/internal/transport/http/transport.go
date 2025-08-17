package http

import (
	"io"
	"mime/multipart"
)

type HttpTransport struct {
}

func NewHttpTransport() *HttpTransport {
	return &HttpTransport{}
}

func nn(io.ReadCloser) {}

func ff() {
	f := multipart.FileHeader{}
	file, _ := f.Open()
	nn(file)
}
