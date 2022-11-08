package demo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"math"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func Get() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:55004",
		Password: "redispw",
		DB:       0,
		Username: "default",
	})
	list := []string{"1a", "2b"}
	rdb.LPush(ctx, "list1", list)
	for i := 0; i < 500; i++ {
		go t()
	}

	time.Sleep(time.Minute)
}

func t() {
	cmd := exec.Command("php", "artisan", "test", "344")
	cmd.Dir = "/Users/ace/php-app/"
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func Tt() {
	s := " 哈哈ss，你好呀你好哈哈当时的哈哈发发发"
	n := strings.Trim(s, "发")
	math.Abs(-1)

	fmt.Println(n)
}

type HandleFunc func(c *Context)

type Context struct {
	index   int8
	handles []HandleFunc
}

func (c *Context) Use(f ...HandleFunc) {
	c.handles = append(c.handles, f...)
}

func (c *Context) Get(f ...HandleFunc) {
	c.handles = append(c.handles, f...)
}

func (c *Context) Next() {
	if c.index < 8 {
		c.index++
		c.handles[c.index](c)
	}
}

func (c *Context) Run() {

	c.handles[0](c)

}

func Mid1() HandleFunc {
	return func(c *Context) {
		fmt.Println("m1")
		c.Next()
	}
}

func Mid2() HandleFunc {
	return func(c *Context) {
		fmt.Println("m2")
		c.Next()
	}
}

func Mid3() HandleFunc {
	return func(c *Context) {
		fmt.Println("m3")
		c.Next()
	}
}

type Person struct {
	Age  int
	Name string
}

func (p Person) Get() int {
	return p.Age
}

func (p *Person) Set(a int) {
	p.Age = a
}

// 初始化pool
var personPool = sync.Pool{
	New: func() interface{} {
		return new(Person)
	},
}

func SyncPool() {
	// 获取一个实例

}
