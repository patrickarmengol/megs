// linked-list deque
package lldeque

import "fmt"

type node[T any] struct {
	data T
	prev *node[T]
	next *node[T]
}

type Deque[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *Deque[T] {
	return &Deque[T]{}
}

func Of[T any](vals ...T) *Deque[T] {
	d := &Deque[T]{}
	d.PushBack(vals...)
	return d
}

func (d *Deque[T]) PushFront(vals ...T) {
	for _, v := range vals {
		n := &node[T]{v, nil, nil}
		if d.IsEmpty() {
			d.head = n
			d.tail = n
		} else {
			d.head.prev = n
			n.next = d.head
			d.head = n
		}
		d.size++
	}
}

func (d *Deque[T]) PushBack(vals ...T) {
	for _, v := range vals {
		n := &node[T]{v, nil, nil}
		if d.IsEmpty() {
			d.head = n
			d.tail = n
		} else {
			d.tail.next = n
			n.prev = d.tail
			d.tail = n
		}
		d.size++
	}
}

func (d *Deque[T]) PopFront() T {
	if d.IsEmpty() {
		panic("can't pop from empty deque")
	}
	res := d.head.data
	// reassign head
	if d.head.next != nil {
		d.head, d.head.next, d.head.next.prev = d.head.next, nil, nil
	}
	d.size--
	return res
}

func (d *Deque[T]) PopBack() T {
	if d.IsEmpty() {
		panic("can't pop from empty deque")
	}
	res := d.tail.data
	if d.tail.prev != nil {
		d.tail, d.tail.prev, d.tail.prev.next = d.tail.prev, nil, nil
	}
	d.size--
	return res
}

func (d *Deque[T]) Front() T {
	if d.IsEmpty() {
		panic("can't pop from empty deque")
	}
	return d.head.data
}

func (d *Deque[T]) Back() T {
	if d.IsEmpty() {
		panic("can't pop from empty deque")
	}
	return d.tail.data
}

func (d *Deque[T]) getNode(index int) *node[T] {
	// TODO: improve by checking if faster from tail
	cur := d.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (d *Deque[T]) Push(index int, vals ...T) {
	if index < 0 || index > d.size {
		panic(fmt.Sprintf("invalid index [%d] for push operation on deque of size %d", index, d.size))
	}

	// TODO: clean this mess up
	var ahead *node[T]
	var behind *node[T]
	if d.IsEmpty() {
		ahead = nil
		behind = nil
	} else if index == d.size {
		ahead = nil
		behind = d.getNode(d.size - 1)

	} else {
		ahead = d.getNode(index)
		behind = ahead.prev
	}

	for _, v := range vals {
		n := &node[T]{v, nil, nil}
		if behind == nil {
			d.head = n
		} else {
			behind.next = n
		}
		n.prev = behind
		behind = n
		d.size++
	}
	behind.next = ahead
	if ahead == nil {
		d.tail = behind
	} else {
		ahead.prev = behind
	}
}

func (d *Deque[T]) Pop(index int) T {
	if d.IsEmpty() {
		panic("can't pop from empty deque")
	}
	if index < 0 || index >= d.size {
		panic(fmt.Sprintf("invalid index [%d] for pop operation on deque of size %d", index, d.size))
	}

	n := d.getNode(index)

	n.prev.next, n.next.prev = n.next, n.prev

	n.prev, n.next = nil, nil

	return n.data
}

func (d *Deque[T]) Set(index int, val T) {
	if d.IsEmpty() {
		panic("can't set value on empty deque")
	}
	if index < 0 || index >= d.size {
		panic(fmt.Sprintf("invalid index [%d] for set operation on deque of size %d", index, d.size))
	}

	d.getNode(index).data = val
}

func (d *Deque[T]) Get(index int) T {
	if d.IsEmpty() {
		panic("can't get value on empty deque")
	}
	if index < 0 || index >= d.size {
		panic(fmt.Sprintf("invalid index [%d] for get operation on deque of size %d", index, d.size))
	}

	return d.getNode(index).data
}

func (d *Deque[T]) Reverse() {
	c := d.head
	for c != nil {
		c.prev, c.next = c.next, c.prev
		c = c.prev
	}
	d.head, d.tail = d.tail, d.head
}

func (d *Deque[T]) IsEmpty() bool {
	return d.Len() == 0
}

func (d *Deque[T]) Len() int {
	return d.size
}

func (d *Deque[T]) Members() []T {
	c := d.head
	s := make([]T, 0, d.Len())
	for c != nil {
		s = append(s, c.data)
		c = c.next
	}
	return s
}

func (d *Deque[T]) String() string {
	return fmt.Sprintf("deque%v", d.Members())
}

func (d *Deque[T]) Clone() *Deque[T] {
	return Of[T](d.Members()...)
}
