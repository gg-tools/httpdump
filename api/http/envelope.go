package http

import (
	"encoding/json"
	"github.com/gg-tools/httpdump/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
)

type envelope struct {
	ErrCode int         `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func getPaging(ctx echo.Context) (page, pageSize int) {
	page = utils.Int(ctx.QueryParam("page"), defaultPage)
	pageSize = utils.Int(ctx.QueryParam("page_size"), defaultPageSize)
	if page <= 0 {
		page = defaultPage
	}
	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	return
}

func jsonOK(ctx echo.Context, data interface{}) error {
	header := ctx.Response().Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	}

	envelope := envelope{
		Data: data,
	}
	jsonStr, err := json.Marshal(envelope)
	if err != nil {
		return jsonError(ctx, err)
	}

	return ctx.String(http.StatusOK, string(jsonStr))
}

type httpError interface {
	error
	HttpCode() int
}

func jsonError(ctx echo.Context, err error) error {
	if err == nil {
		return jsonOK(ctx, nil)
	}

	httpCode := http.StatusInternalServerError
	if httpErr, ok := err.(httpError); ok {
		httpCode = httpErr.HttpCode()
	}

	return ctx.JSON(httpCode, envelope{
		ErrCode: httpCode,
		Msg:     err.Error(),
	})
}
