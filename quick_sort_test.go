package demo

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestQuicksort(t *testing.T) {
	a := []int{1, 3, 5, 6, 9, 1, 3, 5, 67, 4, 0, 111}
	fmt.Println(Quicksort(a))

	x := 311
	fmt.Println(Find(a, x, 0, len(a)-1))
}

func TestNewLock(t *testing.T) {
	var x int
	lock := NewLock()
	for i := 0; i < 100; i++ {
		ii := i
		go func(xx *int, l *Lock, ii int) {
			l.Lock()
			*xx++
			fmt.Println(*xx)
			time.Sleep(time.Second / 10)
			l.UnLock()
		}(&x, lock, ii)
		time.Sleep(time.Second / 10)
		go func(iix int) {
			if lock.TryLock() {
				fmt.Println("get-----" + strconv.Itoa(iix))
				lock.UnLock()
			}
		}(ii)
	}

	time.Sleep(time.Hour)
}
