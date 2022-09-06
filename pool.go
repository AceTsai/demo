package demo

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	f func() error
}

func NewTask (f func() error) *Task {
	t := new(Task)
	t.f = f
	return t
}

func (t *Task) Execute()  {
	time.Sleep(time.Second)
	err := t.f()
	if err != nil {
		return
	}
}

type Pool struct {
	EntryCh chan *Task
	workNum int
}

func NewPool (n int) *Pool {
	return &Pool{
		EntryCh: make(chan *Task),
		workNum: n,
	}
}

func (p *Pool) worker(id int)  {
	for task := range p.EntryCh {
		task.Execute()
		fmt.Println(strconv.Itoa(id) + "finished")
	}
}

func (p *Pool) Run() {
	for i := p.workNum; i > 0; i-- {
		go p.worker(i)
	}
}

func (p *Pool) Close() {
	close(p.EntryCh)
}