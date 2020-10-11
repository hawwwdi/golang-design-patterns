package functionaloption

import (
	"time"
)

type cache struct {
	maxAge     time.Duration
	maxEntries int
}

type option func(*cache) error

func NewCache(options ...option) (*cache, error) {
	c := new(cache)
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func SetMaxAge(t time.Duration) option {
	return func(c *cache) error {
		c.maxAge = t
		return nil
	}
}

func SetMaxEntries(e int) option {
	return func(c *cache) error {
		c.maxEntries = e
		return nil
	}
}
