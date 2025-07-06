package context

import "errors"

type Context struct {
	kv map[string]any
}

func NewContext() *Context {
	return &Context{
		kv: make(map[string]any),
	}
}

func (c *Context) Set(key string, value any) {
	(*c).kv[key] = value
}

func (c Context) Get(key string) any {
	return c.kv[key]
}

func (c *Context) Del(key string) error {
	if _, ok := c.kv[key]; !ok {
		return errors.New("key does not exist")
	}

	delete((*c).kv, key)
	return nil
}

func (c Context) Has(key string) bool {
	_, has := c.kv[key]
	return has
}
