package service

import "github.com/gg-tools/http-dump/model"

type Dumper interface {
	DumpRequest(dumpKey string, request *model.Request) error
	ListRequests(dumpKey string, page, pageSize int) ([]*model.Request, error)
}
