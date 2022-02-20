package high_level_code

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type Cluster struct {
	opts options
}

type options struct {
	connectionTimeout time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	logError          func(ctx context.Context, err error)
}

// 通过一个选项实现为一个函数指针来达到一个目的：设置选项中的数据的状态
// Golang函数指针的用法
type Option func(c *options)

// 设置某个参数的一个具体实现，用到了闭包的用法。
// 不仅仅只是设置而采用闭包的目的是为了更为优化，更好用，对用户更友好
func LogError(f func(ctx context.Context, err error)) Option {
	return func(opts *options) {
		opts.logError = f
	}
}

// 对关键数据变量的赋值采用一个方法来实现而不是直接设置
func ConnectionTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.connectionTimeout = d
	}
}

func WriteTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.writeTimeout = d
	}
}

func ReadTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.readTimeout = d
	}
}

// 构造函数具体实现，传入相关Option，new一个对象并赋值
// 如果参数很多，也不需要传入很多参数，只需要传入opts ...Option即可
// 除了构造函数这个思想之外，还有一个思想，就是我们要善于利用 struct 封装对象方法，然后再 new 一个对象出来
func NewCluster(opts ...Option) *Cluster {
	clusterOpts := options{}
	for _, opt := range opts {
		// 函数指针的赋值调用
		opt(&clusterOpts)
	}

	cluster := new(Cluster)
	cluster.opts = clusterOpts

	return cluster
}

func Construct() {

	// 前期储备，设定相关参数
	commonsOpts := []Option{
		ConnectionTimeout(1 * time.Second),
		ReadTimeout(2 * time.Second),
		WriteTimeout(3 * time.Second),
		LogError(func(ctx context.Context, err error) {
		}),
	}

	mysqlOpts := []Option{
		ConnectionTimeout(4 * time.Second),
		LogError(func(ctx context.Context, err error) {
		}),
		ReadTimeout(5 * time.Second),
		WriteTimeout(6 * time.Second),
	}

	// 终极操作，构造函数
	cluster := NewCluster(commonsOpts...)
	mysqlCluster := NewCluster(mysqlOpts...)

	// 测试验证
	fmt.Println(cluster.opts.connectionTimeout)
	fmt.Println(cluster.opts.writeTimeout)

	fmt.Println(mysqlCluster.opts.writeTimeout)

}
