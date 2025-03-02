package collection

import (
	"errors"
	"net/http"

	"github.com/sebport0/cartero/internal/request"
)

var (
	ErrRequestAlreadyExists = errors.New("request already exists")
	ErrRequestDoesntExists  = errors.New("request doesn't exists")
)

type Collection struct {
	Name string
	reqs map[string]*request.Request
}

func NewCollection(name string, rs map[string]*request.Request) *Collection {
	return &Collection{Name: name, reqs: rs}
}

func (c *Collection) SendRequest(name string) (*http.Response, error) {
	if !c.requestInCollection(name) {
		return nil, ErrRequestDoesntExists
	}
	resp, err := c.reqs[name].Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Collection) AddRequest(r *request.Request) error {
	if c.requestInCollection(r.Name) {
		return ErrRequestAlreadyExists
	}
	c.reqs[r.Name] = r
	return nil
}

func (c *Collection) DeleteRequest(name string) {
	delete(c.reqs, name)
}

func (c *Collection) GetRequest(name string) *request.Request {
	r, exists := c.reqs[name]
	if !exists {
		return nil
	}
	return r
}

func (c *Collection) requestInCollection(name string) bool {
	if _, exists := c.reqs[name]; exists {
		return true
	}
	return false
}
