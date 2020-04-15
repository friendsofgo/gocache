package gocache

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	key1, key2 := "key1", "key2"
	val1, val2 := 1, 2

	cache := New()
	cache.Set(key1, val1)
	cache.Set(key2, val2)

	got := cache.Get(key1).(int)
	if val1 != got {
		t.Fatalf("Get returned unexpected value - expected: %v, got: %v", val1, got)
	}

	got = cache.Get(key2).(int)
	if val2 != got {
		t.Fatalf("Get returned unexpected value - expected: %v, got: %v", val2, got)
	}
}

func BenchmarkCache1(b *testing.B)       { benchmarkCache(1, b) }
func BenchmarkCache10(b *testing.B)      { benchmarkCache(10, b) }
func BenchmarkCache100(b *testing.B)     { benchmarkCache(100, b) }
func BenchmarkCache1000(b *testing.B)    { benchmarkCache(1000, b) }
func BenchmarkCache10000(b *testing.B)   { benchmarkCache(10000, b) }
func BenchmarkCache100000(b *testing.B)  { benchmarkCache(100000, b) }
func BenchmarkCache1000000(b *testing.B) { benchmarkCache(1000000, b) }

func benchmarkCache(items int, b *testing.B) {
	cache := New()
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
	b.ResetTimer()

	got := cache.Get(key)
	if val != got {
		b.Fatalf("Get returned unexpected value - expected: %v, got: %v", val, got)
	}

	//fmt.Println("Alloc:", m2.Alloc-m1.Alloc, "TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc, "HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}