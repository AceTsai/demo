package demo

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Ctx()  {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	for i := 0; i < 100; i++ {
		go process(ctx, i)
	}
	time.Sleep(time.Minute)
}

func process(ctx context.Context, i int) {
	var r string
	select {
		case <- ctx.Done():
			r = "ctxDone"
		case <- time.After(time.Second):
			r = strconv.Itoa(i)
	}
	fmt.Println(r)
}

var s1 string

type s struct {
	Age int
	B string
}
func rangeX () {

	fmt.Println(1 & 0x07)
}

func t(i interface{})  {
	var params []reflect.Value
	params = append(params, reflect.ValueOf(i))
	fmt.Println(reflect.TypeOf(i).Method(0).Func.Call(params))
}

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (c *Cat) Quack() {
	println(c.Name + " meow")
}
