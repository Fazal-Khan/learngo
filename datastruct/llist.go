package datastruct

import (
	"bytes"
	"fmt"
	"sync"
)

type node struct {
	data interface{}
	link *node
}

type llist struct {
	size       int32
	head, tail *node
	lock       sync.Mutex
}

type LList llist

func NewLList() *LList {
	node := new(LList)
	return node
}

func (this *LList) Add(data interface{}) {
	this.add(data, true)
}

func (this *LList) Delete(value interface{}) bool {
	return this.delete(value)
}

func (this *LList) Contains(data interface{}) bool {
	for list := this.head; list != nil; list = list.link {
		if list.data == data {
			return true
		}
	}
	return false
}

func (this *LList) Size() int32 {
	return this.size
}

//--- functions to support Stack & Queue
func (this *LList) AddFirst(data interface{}) {
	this.add(data, false)
}

func (this *LList) RemoveFirst() bool {
	if this.head == nil {
		return false
	}
	return this.Delete(this.head.data)
}

func (this *LList) First() (bool, interface{}) {
	if this.head == nil {
		return false, nil
	}
	return true, this.head.data
}

func (this *LList) Last() (bool, interface{}) {
	if this.tail == nil {
		return false, nil
	}
	return true, this.tail.data
}

//------ util methods -------------------->
//------ toString in golang
func (this *LList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("LList [")
	list := this.head
	for list != nil {
		buffer.WriteString(fmt.Sprint(list.data))
		list = list.link
		if list != nil {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

//------ debug info
func (this *LList) Print() {
	var headdata, taildata interface{}
	if this.head != nil {
		headdata = this.head.data
	}

	if this.tail != nil {
		taildata = this.tail.data
	}
	fmt.Println("Size:", this.Size(), "Head=", headdata, "Tail=", taildata, " --> ", this)
}

//------------ private methods---------------------->
func (this *LList) delete(value interface{}) bool {
	result := false
	if this.head != nil && this.head.data == value {
		this.head = this.head.link
		if this.head == nil {
			this.tail = nil
		}
		result = true
	} else if this.head != nil {
		prev := this.head
		for curr := this.head.link; curr != nil; curr = curr.link {
			if curr.data == value {
				prev.link = curr.link
				if this.tail == curr {
					this.tail = prev
				}
				result = true
				break
			}
		}
	}
	if result {
		this.size -= 1
	}
	return result
}

func (this *LList) add(data interface{}, action bool) {
	lnode := new(node)
	lnode.data = data
	if this.head == nil {
		this.head = lnode
		this.tail = this.head
	} else if action {
		this.tail.link = lnode
		this.tail = lnode
	} else {
		lnode.link = this.head
		this.head = lnode
	}
	this.size += 1
}
