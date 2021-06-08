package service

import (
	"github.com/gg-tools/httpdump/model"
	"github.com/gg-tools/httpdump/model/repository"
	"time"
)

type cacheEntry struct {
	requests []*model.Request
}

func (e *cacheEntry) Add(request *model.Request) {
	e.requests = append(e.requests, request)
}

func (e *cacheEntry) List(page, pageSize int) (requests []*model.Request) {
	offset := pageSize * (page - 1)
	count := 0
	for i := len(e.requests) - offset - 1; i >= 0 && count < pageSize; i-- {
		requests = append(requests, e.requests[i])
		count++
	}
	return
}

type dumper struct {
	cache repository.Cache
}

func NewDumper(cacheRepo repository.Cache) *dumper {
	return &dumper{cacheRepo}
}

func (d *dumper) DumpRequest(dumpKey string, request *model.Request) error {
	request.Time = time.Now()

	entry := d.getEntry(dumpKey)
	entry.Add(request)
	d.setEntry(dumpKey, entry)
	return nil
}

func (d *dumper) ListRequests(dumpKey string, page, pageSize int) ([]*model.Request, error) {
	entry := d.getEntry(dumpKey)
	requests := entry.List(page, pageSize)
	return requests, nil
}

func (d *dumper) getEntry(dumpKey string) *cacheEntry {
	entry := &cacheEntry{}
	if e, ok := d.cache.Get(dumpKey); ok {
		entry = e.(*cacheEntry)
	}

	return entry
}

func (d *dumper) setEntry(dumpKey string, entry *cacheEntry) {
	d.cache.Set(dumpKey, entry, time.Hour)
}
