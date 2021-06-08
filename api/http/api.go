package http

import (
	"github.com/gg-tools/http-dump/model"
	"github.com/gg-tools/http-dump/model/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

type handler struct {
	dumper service.Dumper
}

func NewHandler(dumper service.Dumper) *handler {
	return &handler{dumper}
}

func (h *handler) Route(e *echo.Echo) {
	e.Any("/dumps/:key", h.receiveRequest)

	api := e.Group("/api")
	api.GET("/dumps/:key", h.listRequests)
}

func (h *handler) receiveRequest(c echo.Context) error {
	request := c.Request()
	body := request.Body
	defer body.Close()

	dumpKey := c.Param("key")
	url := request.URL.String()
	method := request.Method
	headers := request.Header
	bodyContent, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	req := &model.Request{
		URL:     url,
		Method:  method,
		Headers: headers,
		Body:    bodyContent,
	}
	return h.dumper.DumpRequest(dumpKey, req)
}

func (h *handler) listRequests(c echo.Context) error {
	dumpKey := c.Param("key")
	page, pageSize := getPaging(c)

	requests, err := h.dumper.ListRequests(dumpKey, page, pageSize)
	if err != nil {
		return err
	}
	return jsonOK(c, adapterRequests(requests))
}
