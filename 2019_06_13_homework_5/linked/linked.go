package linked

import (
	"fmt"
	"strings"
)

type linkedList struct {
	head    *item
	current *item
	tail    *item
}

type item struct {
	value interface{}
	next  *item
	prev  *item
}

func CreateLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) First() *item {
	return l.head
}

func (l *linkedList) Last() *item {
	return l.tail
}

func (l *linkedList) Len() int {
	var count int
	current := l.head

	if current != nil {
		count = 1
	} else {
		return count
	}

	for current.next != nil {
		current = current.next
		count++
	}
	return count
}

func (l *linkedList) PushFront(v interface{}) {

	it := &item{
		value: v,
	}

	if l.head == nil {
		l.head, l.tail, l.current = it, it, it
	} else {
		current := l.head
		it.next = current
		current.prev = it
		l.head = it
	}
}

func (l *linkedList) PushBack(v interface{}) {

	it := &item{
		value: v,
	}

	if l.head == nil {
		l.head, l.tail, l.current = it, it, it
	} else {
		current := l.tail
		it.prev = current
		current.next = it
		l.tail = it
	}
}

func (l *linkedList) Remove(it *item) bool {

	if l.head == it {
		if l.Len() == 1 {
			l.head, l.current, l.tail = nil, nil, nil
			return true
		}
		l.head = it.next
		l.head.prev = nil
		return true
	}

	if l.tail == it {
		l.tail = it.prev
		l.tail.next = nil
		return true
	}

	current := l.head.next

	for current != nil {
		if current == it {
			current.prev.next = current.next
			current.next.prev = current.prev
			return true
		}
		current = current.next
	}
	return false
}

func (l *linkedList) String() string {

	current := l.head

	if current == nil {
		return "LinkedList: {{ }}"
	}

	builder := strings.Builder{}
	builder.WriteString("LinkedList: {{ ")
	builder.WriteString(current.String())

	for current.next != nil {
		current = current.next
		builder.WriteString(" -> ")
		builder.WriteString(current.String())
	}

	builder.WriteString(" }}")
	return builder.String()
}

func (it *item) Value() interface{} {
	return it.value
}

func (it *item) Next() *item {
	return it.next
}

func (it *item) Prev() *item {
	return it.prev
}

func (it *item) String() string {
	return fmt.Sprintf("%v", it.Value())
}
