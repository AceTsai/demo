package demo

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	var a = 1
	var b = 2
	a, b = b, a
	fmt.Println(a, b)
}
