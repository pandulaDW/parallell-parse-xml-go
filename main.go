package main

import (
	"strings"
	"sync"
)

var pool sync.Pool

func init() {
	pool = sync.Pool{
		New: func() interface{} {
			return strings.Builder{}
		},
	}

	// seed the pool with 20 string builders
	for i := 0; i < 20; i++ {
		sb := strings.Builder{}
		pool.Put(sb)
	}
}

func main() {
	concurrentProcessing("lei", "LEI")
	// concurrentProcessing("rr", "Relationship")
}
