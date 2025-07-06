package context

import (
	"errors"
)

type ContextKey string

const (
	PokedexAPI ContextKey = "PokedexAPI"
)

type Context struct {
	kv map[ContextKey]any
}

func NewContext() *Context {
	return &Context{
		kv: make(map[ContextKey]any),
	}
}

func (c *Context) Set(key ContextKey, value any) {
	(*c).kv[key] = value
}

func (c Context) Get(key ContextKey) any {
	return c.kv[key]
}

func (c *Context) Del(key ContextKey) error {
	if _, ok := c.kv[key]; !ok {
		return errors.New("key does not exist")
	}

	delete((*c).kv, key)
	return nil
}

func (c Context) Has(key ContextKey) bool {
	_, has := c.kv[key]
	return has
}
