package id

import (
	"sync/atomic"
)

type Generator struct {
	counter int64
}

func (g *Generator) Next() int64 {
	return atomic.AddInt64(&g.counter, 1)
}

func New() *Generator {
	return &Generator{}
}
