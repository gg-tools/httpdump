package model

import (
	"net/http"
	"time"
)

type Request struct {
	URL     string
	Method  string
	Headers http.Header
	Body    []byte
	Time    time.Time
}
