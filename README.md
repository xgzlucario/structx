# structx

Data structures and algorithms implemented using generics.

Currently, structx provides the following types of data structures to support generic types:

- `List`
- `Map`、`SyncMap`
- `LSet (ListSet)`
- `ZSet (SortedSet)`
- `Pool`
- `Skiplist`
- `Cache`

### List

`List` is a data structure wrapping basic type `slice`.  Compare to basic slice type, List is `sequential`, `sortable`, and `nice wrappered`.

#### usage

```go
ls := structx.NewList(1,2,3)
ls.RPush(4) // [1,2,3,4]
ls.LPop() // 1 [2,3,4]
ls.Reverse() // [4,3,2]

ls.Index(1) // 3
ls.Find(4) // 0

// shift
ls.RShift() // [2,4,3]
ls.Top(1) // [4,2,3]

// Less
ls.SetLess(func(i, j int) bool {
	return l.Index(i) < l.Index(j)
})
// Ascending
ls.SetOrder(true)

ls.Sort() // [2,3,4]
ls.Max() // 4
```

### LSet

`LSet` uses `Map + List` as the storage structure. LSet is Inherited from `List`, where the elements are `sequential` and have `good iterative performance`, as well as `richer api`. When the data volume is small only `list` is used.

#### **usage**

```go
s := structx.NewLSet(1,2,3,4,1) // [1,2,3,4]

s.Remove(3) // [1,2,4]
s.Add(1) // [1,2,4]
s.Add(5) // [1,2,4,5]

// shift
s.Reverse() // [5,4,2,1]
s.Top(2) // [2,5,4,1]
s.Rpop() // [5,4,1]

s.Range(func(k int) bool {
    // do something...
})
newS := structx.NewLSet(1,2,3) // [1,2,3]

union := s.Union(newS) // [0,1,2,3)
intersect := s.Intersect(newS) // [1,2]
diff := s.Difference(newS) // [0,3]
```

#### **Benchmark**

Compare with mapset [deckarep/golang-set](https://github.com/deckarep/golang-set).

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

