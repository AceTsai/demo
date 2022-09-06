package demo

import (
	"fmt"
	"testing"
	"time"
)

func TestLru_Get(t *testing.T) {
	cache := NewLru(3)
	for i:=100; i>0;i-- {
		go func() {
			cache.Set("h1", "v1")
			cache.Set("h2", "v2")
			cache.Set("h3", "v3")
			cache.Get("h1")
			cache.Set("h4", "v4")
			cache.Get("h2")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(cache.Get("h"))
	cache.All()
}

