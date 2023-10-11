package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
}

type Task func() string

func Async(t Task) *FutureResult {
	resultChan := make(chan string, 1)
	f := &FutureResult{
		_,
		resultChan,
	}
	go func() {
		f.ResultChan <- t()
		f.Done.Store(true)
	}()
	return f
}
func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	resultChan := make(chan string, 1)
	f := &FutureResult{
		_,
		resultChan,
	}
	go func() {
		defer close(resultChan)
		f.ResultChan <- t()
		f.Done.Store(true)
	}()
	go func() {
		<-time.After(timeout)
		if !f.Done.Load() {
			close(resultChan)
			f.Done.Store(true)
		}
	}()
	return f
}

func (fResult *FutureResult) Await() string {
	result := <-fResult.ResultChan
	return result
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
	combineResultChan := make(chan string, len(fResults))
	combineDone := atomic.Bool{}

	f := &FutureResult{
		Done:       combineDone,
		ResultChan: combineResultChan,
	}

	var wg *sync.WaitGroup
	wg.Add(len(fResults))

	for _, fResult := range fResults {
		go func(fr *FutureResult) {
			defer wg.Done()
			result := fr.Await()
			f.ResultChan <- result
		}(fResult)
	}

	go func() {
		wg.Wait()
		close(combineResultChan)
		combineDone.Store(true)
	}()

	return f
}
