package demo

import (
	"container/list"
	"fmt"
	"sync"
)

type Lru struct {
	rwl sync.RWMutex
	size int
	l *list.List
	m map[string]*list.Element
}

type KV struct {
	k string
	v string
}

func NewLru(size int) *Lru {
	return &Lru{
		size: size,
		l: list.New(),
		m: make(map[string]*list.Element),
	}
}

func (l *Lru) Set(k, v string) {
	l.rwl.Lock()
	defer l.rwl.Unlock()
	kv := new(KV)
	kv.k = k
	kv.v = v
	if ele, ok := l.m[k]; ok {
		ele.Value = kv
	} else {
		l.m[k] = l.l.PushFront(kv)
	}
	for len(l.m) > l.size {
		b := l.l.Back()
		l.l.Remove(b)
		delete(l.m, b.Value.(*KV).k)
	}
}

func (l *Lru) Get (k string) string {
	l.rwl.RLock()
	defer l.rwl.RUnlock()
	if ele, ok := l.m[k]; ok {
		l.l.MoveToFront(ele)
		return ele.Value.(*KV).v
	}
	return ""
}

func (l *Lru) All() {
	for _,v := range l.m{
		fmt.Println(v.Value.(*KV).v)
	}
}
