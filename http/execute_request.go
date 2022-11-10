package http

import (
	"OwlGramServer/http/types"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func ExecuteRequest(url string, options ...RequestOption) ([]byte, error) {
	var opt types.RequestOptions
	for _, option := range options {
		option.Apply(&opt)
	}
	if (opt.MultiPart != nil) == (opt.Body != nil) && opt.MultiPart != nil {
		return nil, fmt.Errorf("must specify either multipart or body")
	}
	if opt.Method == "" {
		opt.Method = "GET"
	}
	client := http.Client{}
	var body io.Reader
	var multiPartWriter *multipart.Writer
	if opt.MultiPart != nil {
		reader := &bytes.Buffer{}
		multiPartWriter = multipart.NewWriter(reader)
		for k, v := range opt.MultiPart.Data {
			_ = multiPartWriter.WriteField(k, v)
		}
		for k, v := range opt.MultiPart.Files {
			file, err := multiPartWriter.CreateFormFile(k, v.FileName)
			if err != nil {
				return nil, err
			}
			_, err = file.Write(v.Content)
			if err != nil {
				return nil, err
			}
		}
		_ = multiPartWriter.Close()
		body = reader
	} else if opt.Body != nil {
		body = bytes.NewBuffer(opt.Body)
	}
	req, err := http.NewRequest(opt.Method, url, body)
	if err != nil {
		return nil, err
	}
	if opt.Headers != nil {
		for k, v := range opt.Headers {
			req.Header.Set(k, v)
		}
	}
	if opt.BearerToken != "" && req.Header.Get("Authorization") == "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", opt.BearerToken))
	}
	if multiPartWriter != nil {
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	}
	req.Header.Add("Accept-Encoding", "identity")
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(do.Body)
	var buf bytes.Buffer
	_, err = io.Copy(&buf, do.Body)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusCreated && do.StatusCode != http.StatusNoContent {
		opt.Retries--
		if opt.Retries > 0 {
			time.Sleep(time.Millisecond * 250)
			return ExecuteRequest(url, options...)
		}
		return nil, fmt.Errorf("%d, %s", do.StatusCode, buf.Bytes())
	}
	return buf.Bytes(), nil
}
