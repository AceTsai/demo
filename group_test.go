package demo

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	Get()
}

func TestTt(t *testing.T) {
	ctx := new(Context)
	ctx.handles = make([]HandleFunc, 0)

	ctx.Use(Mid1())
	ctx.Use(Mid2())
	ctx.Use(Mid3())

	ctx.Get(func(c *Context) {
		fmt.Println("m33")
	})
	fmt.Println(len(ctx.handles))
	fmt.Println(ctx.index)
	time.Sleep(time.Second)
	ctx.Run()
}

func TestSyncPool(t *testing.T) {
	personPool.Get()
	newPerson := personPool.Get().(*Person)
	newPerson.Set(100)
	fmt.Println(newPerson.Age)
	fmt.Println(newPerson.Get())
}
