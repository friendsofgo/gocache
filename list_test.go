package gocache

import "testing"

func TestList(t *testing.T) {
	t.Run("given an empty list then head and tail are nil", func(t *testing.T) {
		l := new(list)

		if nil != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", nil, l.head)
		}

		if nil != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", nil, l.tail)
		}
	})

	t.Run("given a single item list then and tail are the same item", func(t *testing.T) {
		item1 := &listItem{key: "one"}

		l := new(list)
		l.append(item1)

		if item1 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item1, l.head)
		}

		if item1 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item1, l.tail)
		}
	})

	t.Run("given a list with two items then one corresponds to the head and the other one to the tail", func(t *testing.T) {
		item1, item2 := &listItem{key: "one"}, &listItem{key: "two"}

		l := new(list)
		l.append(item1)
		l.append(item2)

		if item1 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item1, l.head)
		}

		if item2 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item2, l.tail)
		}
	})

	t.Run("given a list with three items then first one corresponds to the head and the last one to the tail", func(t *testing.T) {
		item1, item2, item3 := &listItem{key: "one"}, &listItem{key: "two"}, &listItem{key: "three"}

		l := new(list)
		l.append(item1)
		l.append(item2)
		l.append(item3)

		if item1 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item1, l.head)
		}

		if item3 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item3, l.tail)
		}
	})

	t.Run("given a single item list then the update does nothing", func(t *testing.T) {
		item1 := &listItem{key: "one"}

		l := new(list)
		l.append(item1)

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
		item1, item2 := &listItem{key: "one"}, &listItem{key: "two"}

		l := new(list)
		l.append(item1)
		l.append(item2)

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

		if item1 != item2.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", item1, item2.next)
		}

		if item2 != item1.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", item2, item1.prev)
		}

		if nil != item1.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", nil, item1.next)
		}
	})

	t.Run("given a list with three items then the update sets the second one as the head", func(t *testing.T) {
		item1, item2, item3 := &listItem{key: "one"}, &listItem{key: "two"}, &listItem{key: "three"}

		l := new(list)
		l.append(item1)
		l.append(item2)
		l.append(item3)

		l.update(item2)

		if item2 != l.head {
			t.Fatalf("unexpected list.head - expected: %v, got: %v", item2, l.head)
		}

		if item3 != l.tail {
			t.Fatalf("unexpected list.tail - expected: %v, got: %v", item3, l.tail)
		}

		if nil != item2.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", nil, item2.prev)
		}

		if item1 != item2.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", item1, item2.next)
		}

		if item2 != item1.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", item2, item1.prev)
		}

		if item3 != item1.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", item3, item1.next)
		}

		if item1 != item3.prev {
			t.Fatalf("unexpected item.prev - expected: %v, got: %v", item1, item3.prev)
		}

		if nil != item3.next {
			t.Fatalf("unexpected item.next - expected: %v, got: %v", nil, item3.next)
		}
	})
}
