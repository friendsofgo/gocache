package gocache

type list struct {
	head *listItem
	tail *listItem
}

func (l *list) prepend(item *listItem) {
	if l.head == nil {
		l.head = item
		l.tail = item
		return
	}

	l.head.prepend(item)
	l.head = item
}

func (l *list) update(item *listItem) {
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

func (l *list) pop() (item *listItem) {
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

type listItem struct {
	next *listItem
	prev *listItem
	key  string
	val  interface{}
}

func (item *listItem) prepend(item2 *listItem) {
	item.prev = item2
	item2.next = item
}
