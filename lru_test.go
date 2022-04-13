package gocache

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestLRU(t *testing.T) {
	key1, key2, key3, key4 := "key1", "key2", "key3", "key4"
	val1, val2, val3, val4 := 1, 2, 3, 4

	cache := NewLRU[int](2)
	cache.Set(key1, val1)
	cache.Set(key2, val2)
	cache.Set(key3, val3)
	cache.Set(key4, val4)

	got := cache.Get(key1)
	if 0 != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", nil, got)
	}

	got = cache.Get(key2)
	if 0 != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", nil, got)
	}

	got = cache.Get(key3)
	if val3 != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", val3, got)
	}

	got = cache.Get(key4)
	if val4 != got {
		t.Fatalf("unexpected cache.Get - expected: %v, got: %v", val4, got)
	}
}

func BenchmarkLRU1x10(b *testing.B)       { benchmarkLRU(1, 10, b) }
func BenchmarkLRU10x10(b *testing.B)      { benchmarkLRU(10, 10, b) }
func BenchmarkLRU100x10(b *testing.B)     { benchmarkLRU(100, 10, b) }
func BenchmarkLRU1000x10(b *testing.B)    { benchmarkLRU(1000, 10, b) }
func BenchmarkLRU10000x10(b *testing.B)   { benchmarkLRU(10000, 10, b) }
func BenchmarkLRU100000x10(b *testing.B)  { benchmarkLRU(100000, 10, b) }
func BenchmarkLRU1000000x10(b *testing.B) { benchmarkLRU(1000000, 10, b) }

func BenchmarkLRU1x1000(b *testing.B)       { benchmarkLRU(1, 1000, b) }
func BenchmarkLRU10x1000(b *testing.B)      { benchmarkLRU(10, 1000, b) }
func BenchmarkLRU100x1000(b *testing.B)     { benchmarkLRU(100, 1000, b) }
func BenchmarkLRU1000x1000(b *testing.B)    { benchmarkLRU(1000, 1000, b) }
func BenchmarkLRU10000x1000(b *testing.B)   { benchmarkLRU(10000, 1000, b) }
func BenchmarkLRU100000x1000(b *testing.B)  { benchmarkLRU(100000, 1000, b) }
func BenchmarkLRU1000000x1000(b *testing.B) { benchmarkLRU(1000000, 1000, b) }

func BenchmarkLRU1x100000(b *testing.B)       { benchmarkLRU(1, 100000, b) }
func BenchmarkLRU10x100000(b *testing.B)      { benchmarkLRU(10, 100000, b) }
func BenchmarkLRU100x100000(b *testing.B)     { benchmarkLRU(100, 100000, b) }
func BenchmarkLRU1000x100000(b *testing.B)    { benchmarkLRU(1000, 100000, b) }
func BenchmarkLRU10000x100000(b *testing.B)   { benchmarkLRU(10000, 100000, b) }
func BenchmarkLRU100000x100000(b *testing.B)  { benchmarkLRU(100000, 100000, b) }
func BenchmarkLRU1000000x100000(b *testing.B) { benchmarkLRU(1000000, 100000, b) }

func benchmarkLRU(items, size int, b *testing.B) {
	cache := NewLRU[int](size)
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

	var expected int
	if val >= items-size {
		expected = val
	} else {
		expected = 0
	}
	b.ResetTimer()

	got := cache.Get(key)
	if expected != got {
		b.Fatalf("unexpected cache.Get -  expected: %v, got: %v", expected, got)
	}

	//fmt.Println("Alloc:", m2.Alloc-m1.Alloc, "TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc, "HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}
