package workerpool

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := NewPool(10)
	for i := 0; i < 5; i++ {
	 j := i + 1
	 p.ExecuteRun(func() {
	  fmt.Println("hello ", j)
	 })
	}

	time.Sleep(10 * time.Millisecond)

	fut := p.ExecuteCall(hello)
	res := fut.ResultAwait(30 * time.Millisecond)
	fmt.Println("Result is", res)

	fut1 := p.ExecuteCall(wrap(welcome))
	res1 := fut1.ResultAwait(45 * time.Millisecond)
	fmt.Println("Result for fut1 is:----->>>", res1)

	time.Sleep(1 * time.Second)
	p.Stop()
	fmt.Println("Result after stop for fut is: -->>>>", fut.Result())
	fmt.Println("Done............")

   }

   func hello() interface{} {
	time.Sleep(45 * time.Millisecond)
	return "Hello World."
   }

   func welcome() string {
	time.Sleep(30 * time.Millisecond)
	return "Welcome to go world."
   }

   // wrap is a wrapper function,
   // as func() interface{} signature type does not match func() string signature type,
   // Not sure why :)

   func wrap(f func() string) func() interface{} {
	return func() interface{} { return f() }
   }

