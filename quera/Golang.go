package main

// DO NOT USE ANY IMPORT

type Qutex struct {
	ch chan bool
}

func NewQutex() *Qutex {
	ch1 := make(chan bool, 1)
	return &Qutex{
		ch1,
	}
}

func (q *Qutex) Lock() {
	q.ch <- true
}

func (q *Qutex) Unlock() {
	<-q.ch
}
