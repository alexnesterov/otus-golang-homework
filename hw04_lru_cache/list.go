package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v any) *ListItem
	PushBack(v any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v any) *ListItem {
	item := &ListItem{Value: v, Next: l.head}

	if l.head != nil {
		l.head.Prev = item
	}

	if l.tail == nil {
		l.tail = item
	}

	l.head = item
	l.len++

	return item
}

func (l *list) PushBack(v any) *ListItem {
	item := &ListItem{Value: v, Prev: l.tail}

	if l.tail != nil {
		l.tail.Next = item
	}

	if l.head == nil {
		l.head = item
	}

	l.tail = item
	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	if i == l.head {
		l.head = i.Next
	}

	if i == l.tail {
		l.tail = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
