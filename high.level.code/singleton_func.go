package high_level_code

import "sync"

type config struct {
}

/**
懒汉式，sync.Once是线程安全的，DO方法只会被执行一次
*/
var cfg *config
var oSingle sync.Once

func getInstance() *config {
	oSingle.Do(
		func() {
			cfg = new(config)
		})

	return cfg
}

//饿汉式
func init() {
	cfg = new(config)
}

func newConfig() *config {
	return cfg
}
