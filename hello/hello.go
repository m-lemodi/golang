package main

import "fmt"

type block struct {
	next *block
	prev *block
	key  interface{}
}

type List struct {
	head *block
	tail *block
}

func (L *List) Insert(key interface{}) {
	newer := &block{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = newer
	}

	L.head = newer

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l

}

func (l *List) Display() {
	list := l.head
	for list.next != nil {
		fmt.Printf("%+v -> ", list.key)
		list = list.next
	}
	fmt.Println(list.key)
}

func main() {
	link := List{nil, nil}
	link.Insert(1)
	link.Insert(3)
	link.Insert(5)
	link.Insert(8)
	link.Insert(13)

	link.Display()
}
