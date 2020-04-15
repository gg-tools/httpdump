package http

import (
	"github.com/gg-tools/http-dump/model"
	"net/http"
	"strings"
	"time"
)

type request struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Time    time.Time         `json:"time"`
}

func adapterHeaders(header http.Header) map[string]string {
	headers := make(map[string]string)
	for key, val := range header {
		headers[key] = strings.Join(val, ";")
	}
	return headers
}

func adapterRequest(req *model.Request) *request {
	return &request{
		URL:     req.URL,
		Method:  req.Method,
		Headers: adapterHeaders(req.Headers),
		Body:    string(req.Body),
		Time:    req.Time,
	}
}

func adapterRequests(reqs []*model.Request) (requests []*request) {
	for _, req := range reqs {
		requests = append(requests, adapterRequest(req))
	}
	return
}
