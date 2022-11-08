# structx

Data structures and algorithms implemented using generics.

Currently, structx provides the following types of data structures to support generic types:

- `List`
- `Map`
- `SyncMap`
- `LSet (ListSet)`
- `ZSet (SortedSet)`
- `Pool`
- `Skiplist`
- `Cache`

### LSet

`LSet` is a collection of map + list, has a faster `Range`, `Interset`, `Union` function performance than mapset.

#### **usage**

```go
s := structx.NewLSet[int]()
for i := 0;i < 4;i++ {
    s.Add(i)
}
// (0,1,2,3)

s.Remove(3) // (0,1,2)
s.Add(1) // (0,1,2)
s.Max() // 2

s.Reverse() // (2,1,0)
s.Sort() // (0,1,2)

s.Range(func(k int) {
    // do something...
})
newS := structx.NewLSet(1,2,3) // (1,2,3)

union := s.Union(newS) // (0,1,2,3)
intersect := s.Intersect(newS) // (1,2)
diff := s.Difference(newS) // (0,3)
```

#### **Benchmark**

Compare with mapset [deckarep/golang-set](https://github.com/deckarep/golang-set), **mapsize is 1000**.

```
goos: linux
goarch: amd64
pkg: github.com/xgzlucario/structx/test
cpu: AMD Ryzen 7 5800H with Radeon Graphics  
Benchmark_MapSetRange-16          130693	     8991 ns/op	        0 B/op	      0 allocs/op
Benchmark_LSetRange-16            821851	     1415 ns/op	        0 B/op	      0 allocs/op
Benchmark_MapSetRemove-16      318151948	    3.758 ns/op	        0 B/op	      0 allocs/op
Benchmark_LSetRemove-16        364006822	    3.303 ns/op	        0 B/op	      0 allocs/op
Benchmark_MapSetAdd-16         	   21847	    55064 ns/op	    47871 B/op	     68 allocs/op
Benchmark_LSetAdd-16               17355	    68348 ns/op	    73055 B/op	     78 allocs/op
Benchmark_MapSetUnion-16           12676	    94480 ns/op	    47874 B/op	     68 allocs/op
Benchmark_LSetUnion-16             31516	    38181 ns/op	    30181 B/op	     10 allocs/op
Benchmark_MapSetIntersect-16       14566	    82046 ns/op	    47878 B/op	     68 allocs/op
Benchmark_LSetIntersect-16         37855	    31650 ns/op	    30181 B/op	     10 allocs/op
Benchmark_MapSetDiff-16            30876	    38927 ns/op	     8059 B/op	   1002 allocs/op
Benchmark_LSetDiff-16          	   92643	    12866 ns/op	      153 B/op	      4 allocs/op
```

