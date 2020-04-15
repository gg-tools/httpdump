package repository

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type cache struct {
	c *gocache.Cache
}

func NewCache() *cache {
	c := gocache.New(24*time.Hour, 5*time.Second)
	return &cache{c}
}

func (c *cache) Get(key string) (interface{}, bool) {
	return c.c.Get(key)
}

func (c *cache) Set(key string, val interface{}, expiresIn time.Duration) {
	if val == nil {
		c.c.Delete(key)
		return
	}

	c.c.Set(key, val, expiresIn)
}

func (c *cache) Delete(key string) {
	c.c.Delete(key)
}
