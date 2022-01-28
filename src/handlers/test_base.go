package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type RequestTestCase struct {
	Name            string
	Method          string
	Url             string
	HasAttachment   bool
	FileName        string
	ConvertedMatrix [][]int64
	ExpectedResult  string
	ExpectedStatus  int
	Err             error
}

func BuildRequest(rtc RequestTestCase) (r *http.Request, err error) {
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	if rtc.HasAttachment {
		file, _ := os.Open(rtc.FileName)
		defer file.Close()
		fileWriter, err := multipartWriter.CreateFormFile("file", rtc.FileName)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, err
		}
	}

	multipartWriter.Close()

	request, err := http.NewRequest(rtc.Method, rtc.Url, &requestBody)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return request, nil
}
