package gocache

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestCacheR(t *testing.T) {
	key1, key2, key3, key4 := "key1", "key2", "key3", "key4"
	val1, val2, val3, val4 := 1, 2, 3, 4

	cache := NewR(2)
	cache.Set(key1, val1)
	cache.Set(key2, val2)
	cache.Set(key3, val3)
	cache.Set(key4, val4)

	got := cache.Get(key1)
	if nil != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", nil, got)
	}

	got = cache.Get(key2)
	if nil != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", nil, got)
	}

	got = cache.Get(key3).(int)
	if val3 != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", val3, got)
	}

	got = cache.Get(key4).(int)
	if val4 != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", val4, got)
	}
}

func BenchmarkCacheR1x10(b *testing.B)       { benchmarkCacheR(1, 10, b) }
func BenchmarkCacheR10x10(b *testing.B)      { benchmarkCacheR(10, 10, b) }
func BenchmarkCacheR100x10(b *testing.B)     { benchmarkCacheR(100, 10, b) }
func BenchmarkCacheR1000x10(b *testing.B)    { benchmarkCacheR(1000, 10, b) }
func BenchmarkCacheR10000x10(b *testing.B)   { benchmarkCacheR(10000, 10, b) }
func BenchmarkCacheR100000x10(b *testing.B)  { benchmarkCacheR(100000, 10, b) }
func BenchmarkCacheR1000000x10(b *testing.B) { benchmarkCacheR(1000000, 10, b) }

func BenchmarkCacheR1x1000(b *testing.B)       { benchmarkCacheR(1, 1000, b) }
func BenchmarkCacheR10x1000(b *testing.B)      { benchmarkCacheR(10, 1000, b) }
func BenchmarkCacheR100x1000(b *testing.B)     { benchmarkCacheR(100, 1000, b) }
func BenchmarkCacheR1000x1000(b *testing.B)    { benchmarkCacheR(1000, 1000, b) }
func BenchmarkCacheR10000x1000(b *testing.B)   { benchmarkCacheR(10000, 1000, b) }
func BenchmarkCacheR100000x1000(b *testing.B)  { benchmarkCacheR(100000, 1000, b) }
func BenchmarkCacheR1000000x1000(b *testing.B) { benchmarkCacheR(1000000, 1000, b) }

func BenchmarkCacheR1x100000(b *testing.B)       { benchmarkCacheR(1, 100000, b) }
func BenchmarkCacheR10x100000(b *testing.B)      { benchmarkCacheR(10, 100000, b) }
func BenchmarkCacheR100x100000(b *testing.B)     { benchmarkCacheR(100, 100000, b) }
func BenchmarkCacheR1000x100000(b *testing.B)    { benchmarkCacheR(1000, 100000, b) }
func BenchmarkCacheR10000x100000(b *testing.B)   { benchmarkCacheR(10000, 100000, b) }
func BenchmarkCacheR100000x100000(b *testing.B)  { benchmarkCacheR(100000, 100000, b) }
func BenchmarkCacheR1000000x100000(b *testing.B) { benchmarkCacheR(1000000, 100000, b) }

func benchmarkCacheR(items, size int, b *testing.B) {
	cache := NewR(size)
	var m1, m2 runtime.MemStats

	runtime.ReadMemStats(&m1)
	for i := 0; i < items; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Set(key, i)
	}
	runtime.ReadMemStats(&m2)

	rand.Seed(time.Now().UTC().UnixNano())
	val := rand.Intn(items)
	key := fmt.Sprintf("key%d", val)

	var expected interface{}
	if val >= items-size {
		expected = val
	} else {
		expected = nil
	}
	b.ResetTimer()

	got := cache.Get(key)
	if expected != got {
		b.Fatalf("unexpected cache.Get -  expected: %v, got: %v", expected, got)
	}

	//fmt.Println("Alloc:", m2.Alloc-m1.Alloc, "TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc, "HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}
