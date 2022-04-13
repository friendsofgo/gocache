package gocache

type list[V any] struct {
	head *listItem[V]
	tail *listItem[V]
}

func (l *list[V]) prepend(item *listItem[V]) {
	if l.head == nil {
		l.head = item
		l.tail = item
		return
	}

	l.head.prepend(item)
	l.head = item
}

func (l *list[V]) update(item *listItem[V]) {
	if l.head == item {
		return
	}

	if l.tail == item {
		l.tail = item.prev
	}

	item.prev.next = item.next
	if item.next != nil {
		item.next.prev = item.prev
	}

	item.next = l.head
	l.head.prev = item

	l.head = item
	l.head.prev = nil
}

func (l *list[V]) pop() (item *listItem[V]) {
	if l.head == nil {
		return
	}

	item = l.tail
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	l.tail = item.prev
	return
}

type listItem[V any] struct {
	next *listItem[V]
	prev *listItem[V]
	key  string
	val  V
}

func (item *listItem[V]) prepend(item2 *listItem[V]) {
	item.prev = item2
	item2.next = item
}
