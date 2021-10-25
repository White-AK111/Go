package main

type MyIntQueue struct {
	q []MyInt
}

func NewMyIntQueue() *MyIntQueue {
	return &MyIntQueue{
		q: []MyInt{},
	}
}

func (o *MyIntQueue) Insert(v MyInt) {
	o.q = append(o.q, v)
}
