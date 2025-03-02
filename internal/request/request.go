package request

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Name          string
	Documentation string
	Resp          *http.Response
	req           *http.Request
}

func NewRequest(name string, method, url, doc string, body io.Reader) (*Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return &Request{Name: name, req: req, Documentation: doc}, nil
}

func (r *Request) Do() (*http.Response, error) {
	resp, err := http.DefaultClient.Do(r.req)
	if err != nil {
		return nil, err
	}
	r.Resp = resp
	return r.Resp, nil
}

func (r *Request) URL() *url.URL {
	return r.req.URL
}

func (r *Request) Body() io.ReadCloser {
	return r.req.Body
}

func (r *Request) Method() string {
	return r.req.Method
}

func (r *Request) Header() http.Header {
	return r.req.Header
}

func (r *Request) ContentLength() int64 {
	return r.req.ContentLength
}

func (r *Request) Host() string {
	return r.req.Host
}
