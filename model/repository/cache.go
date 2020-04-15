package repository

import "time"

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{}, expiresIn time.Duration)
	Delete(key string)
}
