package backend

import (
	"errors"
	"net/http"

	"github.com/sebport0/cartero/internal/collection"
	"github.com/sebport0/cartero/internal/request"
)

var (
	ErrCollectionExists       = errors.New("collection with the same name already exists")
	ErrCollectionDoesntExists = errors.New("collection doesn't exist")
)

type Backend struct {
	Collections map[string]*collection.Collection
}

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) SendRequest(colName, reqName string) (*http.Response, error) {
	if !b.collectionExists(colName) {
		return nil, ErrCollectionDoesntExists
	}
	resp, err := b.Collections[colName].SendRequest(reqName)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *Backend) AddCollection(c *collection.Collection) error {
	if b.collectionExists(c.Name) {
		return ErrCollectionExists
	}
	b.Collections[c.Name] = c
	return nil
}

func (b *Backend) DeleteCollection(name string) {
	delete(b.Collections, name)
}

func (b *Backend) AddRequest(collectionName string, r *request.Request) error {
	if !b.collectionExists(collectionName) {
		return ErrCollectionDoesntExists
	}
	b.Collections[collectionName].AddRequest(r)
	return nil
}

func (b *Backend) DeleteRequest(collectionName, reqName string) {
	if !b.collectionExists(collectionName) {
		return
	}
	b.Collections[collectionName].DeleteRequest(reqName)
}

func (b *Backend) collectionExists(name string) bool {
	_, exists := b.Collections[name]
	return exists
}
