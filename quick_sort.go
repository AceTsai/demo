package demo

func Quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := len(a) / 2
	a[pivot], a[right] = a[right], a[pivot]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	Quicksort(a[:left])
	Quicksort(a[left+1:])
	return a
}

func Find(a []int, target, s, e int) int {
	if s > e {
		return -1
	}
	mid := (s + e) / 2
	if target == a[mid] {
		return mid
	}
	if target < a[mid] {
		return Find(a, target, s, mid-1)
	}
	if target > a[mid] {
		return Find(a, target, mid+1, e)
	}
	return -1
}

type Lock struct {
	c chan bool
}

func (l *Lock) Lock() {
	l.c <- true
}

func (l *Lock) UnLock() {
	<-l.c
}

func NewLock() *Lock {
	return &Lock{
		c: make(chan bool, 1),
	}
}

func (l *Lock) TryLock() bool {
	select {
	case l.c <- true:
		return true
	default:
		return false
	}

}
