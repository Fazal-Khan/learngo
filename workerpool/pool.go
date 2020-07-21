//Simple Thread Pool in golang with Simple Future response.

package workerpool

import (
	"fmt"
	"sync"
	"time"
)

// Call is a callable which returns a value
type Call func() interface{}

// Run is a runnable function
type Run func()

// Future is a holder for cal result.
type Future interface {
 Done() bool
 Result() interface{}
 ResultAwait(time.Duration) interface{}
}

// future is an implementation of Future interface
type future struct {
 done    bool
 result  interface{}
 reschan chan struct{}
 sync.RWMutex
}

func (f *future) Done() bool {
 f.RLock()
 defer f.RUnlock()
 return f.done
}

func (f *future) Result() interface{} {
 f.RLock()
 defer f.RUnlock()
 return f.result
}

func (f *future) ResultAwait(d time.Duration) interface{} {
 select {
 case <-f.reschan:
 case <-time.After(d):
 }
 return f.result
}

//--------------------------------------------------------

// worker is worker used in pool
type worker struct {
 name string
}

// start method of worker, called when pool is initialized
func (w *worker) start(runq <-chan Run, quit <-chan struct{}) {
 for {
  select {
  case f := <-runq:
   f()
  case <-quit:
   return
  }
 }
}

//-------------------------------------------------------------

// Pool is pool
type Pool struct {
 pool
}

// pool is a local pool used to implemnt pool methods
type pool struct {
 size int
 runq chan Run
 quit chan struct{}
}

// NewPool is pool constructor
func NewPool(size int) *Pool {
 var p = Pool{
  pool: pool{
   size: size,
   runq: make(chan Run, size),
   quit: make(chan struct{}, size),
  },
 }
 p.init()
 return &p
}

// init method to initialize the pool.
func (p *pool) init() {
 for i := 0; i < p.size; i++ {
  var w = worker{name: fmt.Sprintf("worker %d", i)}
  go w.start(p.runq, p.quit)
 }
}

// Stop the pool after all the submitted tasks are done.
func (p *pool) Stop() {
 p.ExecuteRun(func() {
  close(p.quit)
 })
}

// ExecuteRun runs a runnable object.
func (p *pool) ExecuteRun(run Run) {
 p.runq <- run
}

// ExecuteCall runs a callable function.
func (p *pool) ExecuteCall(call Call) Future {
 var fut = future{reschan: make(chan struct{})}
 fn := func() {
  res := call()
  fut.Lock()
  defer fut.Unlock()
  fut.done = true
  fut.result = res
 }
 p.ExecuteRun(fn)
 return &fut
}