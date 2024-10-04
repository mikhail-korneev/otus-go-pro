package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
		Next:  l.front,
	}

	if l.len == 0 {
		l.front = i
		l.back = i
		l.len++
		return i
	}

	l.front.Prev = i
	l.front = i
	l.len++
	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
		Prev:  l.back,
	}

	if l.len == 0 {
		l.front = i
		l.back = i
		l.len++
		return i
	}

	l.back.Next = i
	l.back = i
	l.len++
	return i
}

func (l *list) Remove(i *ListItem) {
	if i == l.front {
		i.Next.Prev = nil
		l.front = i.Next
		i.Next = nil
		l.len--
		return
	}

	if i == l.back {
		i.Prev.Next = nil
		l.back = i.Prev
		i.Prev = nil
		l.len--
		return
	}

	// item in the middle
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Prev = nil
	i.Next = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	if i == l.back {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		// item in the middle
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.front.Prev = i
	i.Next = l.front
	i.Prev = nil
	l.front = i
}
