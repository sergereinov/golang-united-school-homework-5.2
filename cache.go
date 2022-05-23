package cache

import "time"

type record struct {
	value string
	deadline time.Time
}

type Cache struct {
	cache map[string]record
}

func NewCache() Cache {
	return Cache{make(map[string]record, 0)}
}

func (c *Cache) Get(key string) (string, bool) {
	s, ok := c.cache[key]
	if ok && (s.deadline.IsZero() || s.deadline.After(time.Now())) {
		return s.value, ok
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.cache[key] = record{value, time.Time{}}
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.cache))
	for k, v := range(c.cache) {
		if v.deadline.IsZero() || v.deadline.After(time.Now()) {
			keys = append(keys, k)
		}		
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = record{value, deadline}
}
