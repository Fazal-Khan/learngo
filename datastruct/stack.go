package datastruct

import (
	"fmt"
	"sync"
)

// first in first out
type Stack struct {
	list *LList
}

func NewStack() *Stack {
	return &Stack{list: NewLList()}
}

func (this *Stack) Push(data interface{}) {
	this.list.AddFirst(data)
}

func (this *Stack) Pop() (bool, interface{}) {
	ok, data := this.list.First()
	if ok {
		this.list.RemoveFirst()
	}
	return ok, data
}

func (this *Stack) Peek() (bool, interface{}) {
	return this.list.First()
}

//------ util methods -------------------->
func (this *Stack) Lock() sync.Mutex {
	return this.list.lock
}

func (this *Stack) String() string {
	return this.list.String()
}

func (this *Stack) Print() {
	fmt.Println(this)
}
