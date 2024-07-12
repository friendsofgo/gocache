package gocache

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
)

func TestBoundless(t *testing.T) {
	key1, key2 := "key1", "key2"
	val1, val2 := 1, 2

	cache := NewBoundless[string, int]()
	cache.Set(key1, val1)
	cache.Set(key2, val2)

	got := cache.Get(key1)
	if val1 != got {
		t.Fatalf("Get returned unexpected value - expected: %v, got: %v", val1, got)
	}

	got = cache.Get(key2)
	if val2 != got {
		t.Fatalf("Get returned unexpected value - expected: %v, got: %v", val2, got)
	}
}

func BenchmarkBoundless1(b *testing.B)       { benchmarkBoundless(1, b) }
func BenchmarkBoundless10(b *testing.B)      { benchmarkBoundless(10, b) }
func BenchmarkBoundless100(b *testing.B)     { benchmarkBoundless(100, b) }
func BenchmarkBoundless1000(b *testing.B)    { benchmarkBoundless(1000, b) }
func BenchmarkBoundless10000(b *testing.B)   { benchmarkBoundless(10000, b) }
func BenchmarkBoundless100000(b *testing.B)  { benchmarkBoundless(100000, b) }
func BenchmarkBoundless1000000(b *testing.B) { benchmarkBoundless(1000000, b) }

func benchmarkBoundless(items int, b *testing.B) {
	cache := NewBoundless[string, int]()
	var m1, m2 runtime.MemStats

	runtime.ReadMemStats(&m1)
	for i := 0; i < items; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Set(key, i)
	}
	runtime.ReadMemStats(&m2)

	val := rand.Intn(items)
	key := fmt.Sprintf("key%d", val)
	b.ResetTimer()

	got := cache.Get(key)
	if val != got {
		b.Fatalf("Get returned unexpected value - expected: %v, got: %v", val, got)
	}

	//fmt.Println("Alloc:", m2.Alloc-m1.Alloc, "TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc, "HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}
