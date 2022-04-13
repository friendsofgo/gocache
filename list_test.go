package gocache

import "testing"

func TestList(t *testing.T) {
	t.Run("given an empty list then head and tail are nil", func(t *testing.T) {
		l := new(list[int])

		if nil != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", nil, l.head)
		}

		if nil != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", nil, l.tail)
		}
	})

	t.Run("given a single item list then and tail are the same item", func(t *testing.T) {
		item1 := &listItem[int]{key: "one"}

		l := new(list[int])
		l.prepend(item1)

		if item1 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item1, l.head)
		}

		if item1 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item1, l.tail)
		}
	})

	t.Run("given a list with two items then one corresponds to the head and the other one to the tail", func(t *testing.T) {
		item1, item2 := &listItem[int]{key: "one"}, &listItem[int]{key: "two"}

		l := new(list[int])
		l.prepend(item1)
		l.prepend(item2)

		if item2 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item2, l.head)
		}

		if item1 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item1, l.tail)
		}
	})

	t.Run("given a list with three items then first one corresponds to the head and the last one to the tail", func(t *testing.T) {
		item1, item2, item3 := &listItem[int]{key: "one"}, &listItem[int]{key: "two"}, &listItem[int]{key: "three"}

		l := new(list[int])
		l.prepend(item1)
		l.prepend(item2)
		l.prepend(item3)

		if item3 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item3, l.head)
		}

		if item1 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item1, l.tail)
		}
	})

	t.Run("given a single item list then the update does nothing", func(t *testing.T) {
		item1 := &listItem[int]{key: "one"}

		l := new(list[int])
		l.prepend(item1)

		l.update(item1)

		if item1 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item1, l.head)
		}

		if item1 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item1, l.tail)
		}

		if nil != item1.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", nil, item1.prev)
		}

		if nil != item1.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", nil, item1.next)
		}
	})

	t.Run("given a list with two items then the update sets the second one as the head", func(t *testing.T) {
		item1, item2 := &listItem[int]{key: "one"}, &listItem[int]{key: "two"}

		l := new(list[int])
		l.prepend(item1)
		l.prepend(item2)

		l.update(item1)

		if item1 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item1, l.head)
		}

		if item2 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item2, l.tail)
		}

		if nil != item1.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", nil, item1.prev)
		}

		if item2 != item1.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", item2, item1.next)
		}

		if item1 != item2.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", item1, item2.prev)
		}

		if nil != item2.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", nil, item2.next)
		}
	})

	t.Run("given a list with three items then the update sets the second one as the head", func(t *testing.T) {
		item1, item2, item3 := &listItem[int]{key: "one"}, &listItem[int]{key: "two"}, &listItem[int]{key: "three"}

		l := new(list[int])
		l.prepend(item1)
		l.prepend(item2)
		l.prepend(item3)

		l.update(item2)

		if item2 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item2, l.head)
		}

		if item1 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item1, l.tail)
		}

		if nil != item2.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", nil, item2.prev)
		}

		if item3 != item2.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", item3, item2.next)
		}

		if item2 != item3.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", item2, item3.prev)
		}

		if item1 != item3.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", item1, item3.next)
		}

		if item3 != item1.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", item3, item1.prev)
		}

		if nil != item1.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", nil, item1.next)
		}
	})

	t.Run("given an empty list then pop returns nil", func(t *testing.T) {
		l := new(list[int])

		got := l.pop()
		if nil != got {
			t.Fatalf("unexpected list.pop - expected: %v, got: %v", nil, got)
		}

		if nil != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", nil, l.head)
		}

		if nil != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", nil, l.tail)
		}
	})

	t.Run("given a single item list then pop returns the item and the list is empty", func(t *testing.T) {
		item1 := &listItem[int]{key: "one"}

		l := new(list[int])
		l.prepend(item1)

		got := l.pop()
		if item1 != got {
			t.Fatalf("unexpected list.pop - expected: %v, got: %v", item1, got)
		}

		if nil != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", nil, l.head)
		}

		if nil != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", nil, l.tail)
		}
	})

	t.Run("given a list with two items then pop returns the last item and the list is a single one", func(t *testing.T) {
		item1, item2 := &listItem[int]{key: "one"}, &listItem[int]{key: "two"}

		l := new(list[int])
		l.prepend(item1)
		l.prepend(item2)

		got := l.pop()
		if item1 != got {
			t.Fatalf("unexpected list.pop - expected: %v, got: %v", item1, got)
		}

		if item2 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item2, l.head)
		}

		if item2 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item2, l.tail)
		}
	})

	t.Run("given a list with three items then pop returns the last item and the list size is two", func(t *testing.T) {
		item1, item2, item3 := &listItem[int]{key: "one"}, &listItem[int]{key: "two"}, &listItem[int]{key: "three"}

		l := new(list[int])
		l.prepend(item1)
		l.prepend(item2)
		l.prepend(item3)

		got := l.pop()
		if item1 != got {
			t.Fatalf("unexpected list.pop - expected: %v, got: %v", item1, got)
		}

		if item3 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item3, l.head)
		}

		if item2 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item2, l.tail)
		}
	})
}
