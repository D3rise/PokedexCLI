package context

import (
	"errors"
	"sync"
)

type ContextKey string

type Context struct {
	mu sync.RWMutex
	kv map[ContextKey]any
}

func NewContext() *Context {
	return &Context{
		kv: make(map[ContextKey]any),
	}
}

func (c *Context) Set(key ContextKey, value any) {
	(*c).mu.Lock()
	defer (*c).mu.Unlock()

	(*c).kv[key] = value
}

func (c *Context) Get(key ContextKey) any {
	(*c).mu.RLock()
	defer (*c).mu.RUnlock()

	return c.kv[key]
}

func (c *Context) Del(key ContextKey) error {
	(*c).mu.Lock()
	defer (*c).mu.Unlock()

	if _, ok := c.kv[key]; !ok {
		return errors.New("key does not exist")
	}

	delete((*c).kv, key)
	return nil
}

func (c *Context) Has(key ContextKey) bool {
	(*c).mu.RLock()
	defer (*c).mu.RUnlock()

	_, has := c.kv[key]
	return has
}
