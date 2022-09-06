package demo

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {

}

func TestNewTask(t *testing.T) {

}

func TestPool_Run(t *testing.T) {
	pool := NewPool(3)
	fun := func() error {
		rand.Seed(time.Now().Unix())
		time.Sleep(time.Millisecond)
		s := rand.Intn(100)
		fmt.Println("结果" + strconv.Itoa(s))
		return nil
	}
	task := NewTask(fun)
	pool.Run()
	pool.EntryCh <- task
	pool.EntryCh <- task
	pool.EntryCh <- task
	pool.EntryCh <- task
	pool.EntryCh <- task
	pool.EntryCh <- task
	time.Sleep(time.Minute)

}

