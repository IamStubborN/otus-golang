package main

import (
	"fmt"
	"math/rand"
	"time"
)

const fnsNum = 100
const parallelFns = 5
const maxErrorCount = 3

var stop = make(chan struct{}, 1)
var routines = make(chan struct{}, parallelFns)
var errs = make(chan error, maxErrorCount)

var fns = make([]func() error, fnsNum)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < fnsNum; i++ {
		fns[i] = workFunc(i)
	}

	run(fns, parallelFns, maxErrorCount)
}

func workFunc(idx int) func() error {
	return func() error {

		defer func() {
			<-routines
			if len(routines) == 0 {
				stop <- struct{}{}
			}
			if err := recover(); err != nil {
				return
			}
		}()

		if rand.Intn(100) > 95 {
			return fmt.Errorf("Hello. I'm №%d error\n", idx)
		}

		time.Sleep(500 * time.Millisecond) // Для наглядности
		fmt.Printf("Hello. I'm №%d func\n", idx)

		return nil
	}
}

func run(fns []func() error, parallelFns, maxErrorCount int) {
OUTER:
	for _, fn := range fns {
		routines <- struct{}{}
		select {
		case <-stop:
			stop <- struct{}{}
			break OUTER
		default:
			go func(fn func() error) {

				defer func() {
					if err := recover(); err != nil {
						return
					}
				}()

				err := fn()
				if err != nil {
					errs <- err
				}
				if len(errs) == maxErrorCount {
					stop <- struct{}{}
				}
			}(fn)
		}
	}

	if _, ok := <-stop; ok {
		close(stop)
		close(errs)
		close(routines)
	}

	time.Sleep(10 * time.Millisecond) // Для правильного отступа

	fmt.Println("--------------------")
	for val := range errs {
		fmt.Print(val)
	}
}
