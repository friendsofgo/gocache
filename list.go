package gocache

type list struct {
	head *listItem
	tail *listItem
}

func (l *list) append(item *listItem) {
	if l.head == nil {
		l.head = item
		l.tail = item
		return
	}

	l.tail.append(item)
	l.tail = item
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

	//
	//item.append(l.head)
	//l.head = item
	//
	//if l.head == nil {
	//	l.head = item
	//	l.tail = item
	//	return
	//}
	//
	//l.tail.append(item)
	//l.tail = item
}

type listItem struct {
	next *listItem
	prev *listItem
	key  string
	val  interface{}
}
