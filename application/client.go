package application

import (
	"github.com/zuzuleinen/translator/values"
)

const StatusOK = "ok"

type Storage interface {
	Store(req values.StoreRequest)
	Find(req values.GetRequest) string
}

type Client struct {
	storage Storage
}

type Response struct {
	Body string
}

func NewClient(storage Storage) *Client {
	return &Client{storage}
}

func (c *Client) StoreRequest(req values.StoreRequest) Response {
	c.storage.Store(req)
	return Response{Body: StatusOK}
}

func (c *Client) GetRequest(req values.GetRequest) Response {
	return Response{Body: c.storage.Find(req)}
}
