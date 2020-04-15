## Gocache

Gocache is an in-memory cache implementation written in Go.

### TODOs

- [ ] Memory profiling + improvements.
- [ ] Badges (+ code report + unit tests).
- [ ] Concurrent safety.
- [ ] Scan parser method.

### Architectural decisions

Due to the non-existence of generics in Go, the `interface{}` was used instead, what forces us to use type casting
(so reflection) on data fetches.

### Data structures & algorithms

#### Unlimited cache

Unlimited cache that uses a HashMap (`map[string]interface{}`) under the hood.

Time cost analysis:

- Get (`O(1)`) = HashMap lookup (`O(1)`).
- Set (`O(1)`) = HashMap insert (`O(1)`).

#### With replacement cache

Limited (with [LRU replacement](https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)) cache
that uses a HashMap(`map[string]interface{}`) and a doubly-linked list under the hood.

Time cost analysis:

- Get (`O(1)`) = HashMap lookup (`O(1)`) + LinkedList relocation (`O(1)`).
- Set (`O(1)`) = HashMap insert (`O(1)`) + LinkedList insert (`O(1)`) + eviction (`O(1)`).
- eviction (`O(1)`) = HashMap deletion (`O(1)`) + LinkedList deletion (`O(1)`).

### Usage

#### Unlimited cache

Just initialize a new cache and start using it through the `Get` and `Set` methods.

```go
package main

import "github.com/friendsofgo/gocache"

func main() {
	key1, key2 := "key1", "key2"
	val1, val2 := 1, 2

	cache := gocache.New()
	cache.Set(key1, val1)
	cache.Set(key2, val2)

	cache.Get(key1) // returns: 1
	cache.Get(key2) // returns: 2
}
```

#### With replacement cache

Just initialize a new cache with an initial size and start using it through the `Get` and `Set` methods.

```go
package main

import "github.com/friendsofgo/gocache"

func main() {
	key1, key2 := "key1", "key2"
	val1, val2 := 1, 2
    size := 1

	cache := gocache.NewR(size)
	cache.Set(key1, val1)
	cache.Set(key2, val2)

	cache.Get(key1) // returns: nil (eviction)
	cache.Get(key2) // returns: 2
}
```

### Benchmarking

#### Time consumption - O(1)

As you you can see on the attached benchmark logs, the lookup time cost of both implementations corresponds to O(1).

```
BenchmarkCache1-8                 	1000000000	         0.000000 ns/op
BenchmarkCache10-8                	1000000000	         0.000000 ns/op
BenchmarkCache100-8               	1000000000	         0.000000 ns/op
BenchmarkCache1000-8              	1000000000	         0.000000 ns/op
BenchmarkCache10000-8             	1000000000	         0.000000 ns/op
BenchmarkCache100000-8            	1000000000	         0.000001 ns/op
BenchmarkCache1000000-8           	1000000000	         0.000001 ns/op
BenchmarkCacheR1x10-8             	1000000000	         0.000000 ns/op
BenchmarkCacheR10x10-8            	1000000000	         0.000001 ns/op
BenchmarkCacheR100x10-8           	1000000000	         0.000001 ns/op
BenchmarkCacheR1000x10-8          	1000000000	         0.000000 ns/op
BenchmarkCacheR10000x10-8         	1000000000	         0.000000 ns/op
BenchmarkCacheR100000x10-8        	1000000000	         0.000000 ns/op
BenchmarkCacheR1000000x10-8       	1000000000	         0.000001 ns/op
BenchmarkCacheR1x1000-8           	1000000000	         0.000000 ns/op
BenchmarkCacheR10x1000-8          	1000000000	         0.000001 ns/op
BenchmarkCacheR100x1000-8         	1000000000	         0.000000 ns/op
BenchmarkCacheR1000x1000-8        	1000000000	         0.000001 ns/op
BenchmarkCacheR10000x1000-8       	1000000000	         0.000001 ns/op
BenchmarkCacheR100000x1000-8      	1000000000	         0.000001 ns/op
BenchmarkCacheR1000000x1000-8     	1000000000	         0.000001 ns/op
BenchmarkCacheR1x100000-8         	1000000000	         0.000000 ns/op
BenchmarkCacheR10x100000-8        	1000000000	         0.000000 ns/op
BenchmarkCacheR100x100000-8       	1000000000	         0.000001 ns/op
BenchmarkCacheR1000x100000-8      	1000000000	         0.000001 ns/op
BenchmarkCacheR10000x100000-8     	1000000000	         0.000002 ns/op
BenchmarkCacheR100000x100000-8    	1000000000	         0.000001 ns/op
BenchmarkCacheR1000000x100000-8   	1000000000	         0.000001 ns/op
```

#### Memory footprint

![memory-footprint](images/memory_footprint_v1.png)
