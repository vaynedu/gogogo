package async

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"
)

// sentinal errors
var (
	ErrAsyncVariableNil = errors.New("variable is nil")
	ErrAsyncTimeout     = errors.New("timeout")
)

// Async Encapsulation goroutine
type Async struct {
	fs    []func() // 函数
	wg    *sync.WaitGroup
	start time.Time
	cost  time.Duration
	log   *log.Logger
}

func New() *Async {
	return &Async{
		wg:    &sync.WaitGroup{},
		start: time.Now(),
		log:   log.New(os.Stdout, "async:", log.Default().Flags()),
	}
}

// Add add funcnames
func (as *Async) Add(f func()) error {
	if as == nil {
		return ErrAsyncVariableNil
	}
	as.fs = append(as.fs, f)
	return nil
}

// RunAndWait do funcnames
func (as *Async) RunAndWait() error {
	if as == nil {
		return ErrAsyncVariableNil
	}
	for _, fs := range as.fs {
		as.wg.Add(1)
		go func(f func()) {
			defer as.wg.Done()
			f()
		}(fs)

	}
	as.wg.Wait()
	as.cost = time.Since(as.start)
	funcCounts, _ := as.GetFuncCounts()
	as.log.Printf("cost_time:%v,funcs:%d", as.cost, funcCounts)
	return nil
}

func (as *Async) GetFuncCounts() (int, error) {
	if as == nil {
		return 0, ErrAsyncVariableNil
	}
	return len(as.fs), nil
}
