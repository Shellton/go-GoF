package singleton

import "sync"

type Interface interface {
	Function(args []string) string
}

type singleton struct {
}

var hungry Interface = &singleton{}

func HungryMode() Interface {
	return hungry
}

func (*singleton) Function(args []string) string {
	return "result"
}

var lazy Interface
var lazyOnce sync.Once

func LazyMode() Interface {
	lazyOnce.Do(func() {
		lazy = &singleton{}
	})
	return lazy
}
