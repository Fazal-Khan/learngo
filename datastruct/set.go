package datastruct

import (
	"fmt"
)

type Set interface {
	
	Add(data interface{})

	Delete(data interface{}) bool

	Contains(data interface{}) bool

	Size() int

	ToArray() []interface{}

	Print()
}

type set struct {
	mymap map[interface{}]struct{}
	empty struct{}
}

func NewSet() Set {
	return &set{mymap: make(map[interface{}]struct{})}
}

func (this *set) Add(data interface{}) {
	this.mymap[data] = this.empty
}

func (this *set) Delete(data interface{}) bool {
	_, ok := this.mymap[data]
	if ok {
		delete(this.mymap, data)
	}
	return ok
}

func (this *set) Contains(data interface{}) bool {
	_, ok := this.mymap[data]
	return ok
}

func (this *set) Size() int {
	return len(this.mymap)
}

func (this *set) ToArray() []interface{} {
	size := len(this.mymap)
	arr := make([]interface{}, size)
	i := 0
	for k, _ := range this.mymap {
		arr[i] = k
		i++
	}
	return arr
}

//---- just to debug-----
func (this *set) Print() {
	for k, v := range this.mymap {
		fmt.Println(k, v)
	}
}
