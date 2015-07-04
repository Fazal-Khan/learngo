package datastruct

import (
	"fmt"
	"sync"
)

// first in first out
type Queue struct {
	list *LList
}

func NewQueue() *Queue {
	return &Queue{list: NewLList()}
}

func (this *Queue) Push(data interface{}) {
	this.list.Add(data)
}

func (this *Queue) Pop() (bool, interface{}) {
	ok, data := this.list.First()
	if ok {
		this.list.RemoveFirst()
	}
	return ok, data
}

func (this *Queue) Peek() (bool, interface{}) {
	return this.list.First()
}

func (this *Queue) hasNext() bool {
	ok, _ := this.list.First()
	return ok
}

//------ util methods -------------------->
func (this *Queue) Lock() sync.Mutex {
	return this.list.lock
}

func (this *Queue) String() string {
	return this.list.String()
}

func (this *Queue) Print() {
	fmt.Println(this)
}
