package main

import (
	"sync"
)

type CaseChanFan []CaseChan

func (pip CaseChan) Collect(done <-chan bool, num int) (result []*Case) {
	for i := 0; i < num; i++ {
		select {
		case val := <-pip:
			result = append(result, val)
		case <-done:
			return
		}
	}
	return
}
func (pip CaseChan) Fan(done <-chan bool, workers int, fn func(*Case) *Case) CaseChan {
	return pip.FanOut(done, workers, fn).FanIn(done)
}
func (pip CaseChan) FanOut(done <-chan bool, workers int, fn func(*Case) *Case) CaseChanFan {
	fan := CaseChanFan{}
	for i := 0; i < workers; i++ {
		fan = append(fan, pip.worker(done, fn))
	}
	return fan
}
func (pip CaseChan) Filter(done <-chan bool, fn func(*Case) bool) CaseChan {
	out := make(chan *Case)
	go func() {
		defer close(out)
		for val := range pip {
			if fn(val) {
				select {
				case out <- val:
				case <-done:
					return
				}
			}
		}
	}()
	return out
}
func (pip CaseChan) Pipe(done <-chan bool, fn func(*Case) *Case) CaseChan {
	out := make(chan *Case)
	go func() {
		defer close(out)
		for val := range pip {
			select {
			case out <- fn(val):
			case <-done:
				return
			}
		}
	}()
	return out
}
func (pip CaseChan) worker(done <-chan bool, fn func(*Case) *Case) CaseChan {
	out := make(chan *Case)
	go func() {
		defer close(out)
		for val := range pip {
			select {
			case out <- fn(val):
			case <-done:
				return
			}
		}
	}()
	return out
}
func (fan CaseChanFan) FanIn(done <-chan bool) CaseChan {
	var wg sync.WaitGroup
	out := make(chan *Case)
	output := func(pl CaseChan) {
		defer wg.Done()
		for val := range pl {
			select {
			case out <- val:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(fan))
	for _, val := range fan {
		go output(val)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func (fan CaseChanFan) Filter(done <-chan bool, fn func(*Case) bool) (result CaseChanFan) {
	for _, pipe := range fan {
		result = append(result, pipe.Filter(done, fn))
	}
	return
}
func (fan CaseChanFan) Pipe(done <-chan bool, fn func(*Case) *Case) (result CaseChanFan) {
	for _, pipe := range fan {
		result = append(result, pipe.Pipe(done, fn))
	}
	return
}
