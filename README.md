# structx

Data structures and algorithms implemented using generics.

Currently, structx provides the following types of data structures to support generic types：**List**, **Map**,  **SyncMap**, **ZSet** (SortedSet), **LSet** (ListSet).

### LSet

LSet is a collection of map + list, has a faster Range, Interset, Union function performance than mapset.

#### **usage**

```go
s := structx.NewLSet[int]()
for i := 0;i < 4;i++ {
    s.Add(i)
}
// (0,1,2,3)

s.Remove(3) // (0,1,2)
s.Add(1) // (0,1,2)
s.Range(func(k int) {
    // do something...
})
newS := structx.NewLSet(1,2,3) // (1,2,3)

unionRes := s.Union(newS) // (0,1,2,3)
intersectRes := s.Intersect(newS) // (1,2)
```

#### **Benchmark**

Compare with mapset [deckarep/golang-set](https://github.com/deckarep/golang-set), **mapsize is 1000**.

```
goos: linux
goarch: amd64
pkg: github.com/xgzlucario/structx/test
cpu: AMD Ryzen 7 5800H with Radeon Graphics  
Benchmark_MapSetRange-16         133394	     9016 ns/op	       0 B/op	    0 allocs/op
Benchmark_LSetRange-16           773064	     1466 ns/op	       0 B/op	    0 allocs/op
Benchmark_MapSetRemove-16     279567439	    4.300 ns/op	       0 B/op	    0 allocs/op
Benchmark_LSetRemove-16       356212938	    3.416 ns/op	       0 B/op	    0 allocs/op
Benchmark_MapSetAdd-16            21336	    56803 ns/op	   47866 B/op	   68 allocs/op
Benchmark_LSetAdd-16              17331	    69128 ns/op	   73054 B/op	   78 allocs/op
Benchmark_MapSetUnion-16          12068	    97159 ns/op	   47874 B/op	   68 allocs/op
Benchmark_LSetUnion-16            31046	    38477 ns/op	   30181 B/op	   10 allocs/op
Benchmark_MapSetIntersect-16      14470	    82726 ns/op	   47878 B/op	   68 allocs/op
Benchmark_LSetIntersect-16        29746	    40416 ns/op	   30182 B/op	   10 allocs/op
```

