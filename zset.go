package structx

import (
	"fmt"
	"math/rand"
)

const zSkiplistMaxlevel = 32

type (
	skipListLevel[K comparable, S Value] struct {
		forward *skipListNode[K, S]
		span    uint64
	}

	skipListNode[K comparable, S Value] struct {
		key      K
		score    S
		backward *skipListNode[K, S]
		level    []*skipListLevel[K, S]
	}

	node[K comparable, S Value] struct {
		score S
		data  K
	}

	skipList[K comparable, S Value] struct {
		header *skipListNode[K, S]
		tail   *skipListNode[K, S]
		length int64
		level  int
	}

	ZSet[K comparable, S Value] struct {
		dict map[K]*node[K, S]
		zsl  *skipList[K, S]
	}
)

func zslCreateNode[K comparable, S Value](level int, score S, key K) *skipListNode[K, S] {
	n := &skipListNode[K, S]{
		key:   key,
		score: score,
		level: make([]*skipListLevel[K, S], level),
	}
	for i := range n.level {
		n.level[i] = new(skipListLevel[K, S])
	}
	return n
}

func zslCreate[K comparable, S Value]() *skipList[K, S] {
	var key K
	var score S
	return &skipList[K, S]{
		level:  1,
		header: zslCreateNode(zSkiplistMaxlevel, score, key),
	}
}

// Skiplist P = 1/4
const zSkiplistP = 0.25

/*
Returns a random level for the new skiplist node we are going to create.
The return value of this function is between 1 and _ZSKIPLIST_MAXLEVEL
(both inclusive), with a powerlaw-alike distribution where higher
levels are less likely to be returned.
*/
func randomLevel() int {
	level := 1
	for float32(rand.Int31()&0xFFFF) < (zSkiplistP * 0xFFFF) {
		level++
	}
	if level < zSkiplistMaxlevel {
		return level
	}
	return zSkiplistMaxlevel
}

/*
zslInsert a new node in the skiplist. Assumes the element does not already
exist (up to the caller to enforce that). The skiplist takes ownership
of the passed SDS string 'node'.
*/
func (zsl *skipList[K, S]) zslInsert(score S, key K) *skipListNode[K, S] {
	update := make([]*skipListNode[K, S], zSkiplistMaxlevel)
	rank := make([]uint64, zSkiplistMaxlevel)

	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		/* store rank that is crossed to reach the insert position */
		if i == zsl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		if x.level[i] != nil {
			for x.level[i].forward != nil &&
				(x.level[i].forward.score < score ||
					(x.level[i].forward.score == score)) {
				rank[i] += x.level[i].span
				x = x.level[i].forward
			}
		}
		update[i] = x
	}

	/* we assume the element is not already inside, since we allow duplicated
	 * scores, reinserting the same element should never happen since the
	 * caller of zslInsert() should test in the hash table if the element is
	 * already inside or not. */
	level := randomLevel()
	if level > zsl.level {
		for i := zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = uint64(zsl.length)
		}
		zsl.level = level
	}

	x = zslCreateNode(level, score, key)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		/* update span covered by update[i] as x is inserted here */
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	/* increment span for untouched levels */
	for i := level; i < zsl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == zsl.header {
		x.backward = nil
	} else {
		x.backward = update[0]

	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		zsl.tail = x
	}
	zsl.length++
	return x
}

/* Internal function used by zslDelete, zslDeleteByScore and zslDeleteByRank */
func (zsl *skipList[K, S]) zslDeleteNode(x *skipListNode[K, S], update []*skipListNode[K, S]) {
	for i := 0; i < zsl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		zsl.tail = x.backward
	}
	for zsl.level > 1 && zsl.header.level[zsl.level-1].forward == nil {
		zsl.level--
	}
	zsl.length--
}

/* Delete an element with matching score/element from the skiplist.
 * The function returns 1 if the node was found and deleted, otherwise
 * 0 is returned.
 *
 * If 'node' is NULL the deleted node is freed by zslFreeNode(), otherwise
 * it is not freed (but just unlinked) and *node is set to the node pointer,
 * so that it is possible for the caller to reuse the node (including the
 * referenced SDS string at node->node). */
func (zsl *skipList[K, S]) zslDelete(score S, key K) bool {
	update := make([]*skipListNode[K, S], zSkiplistMaxlevel)
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	/* We may have multiple elements with the same score, what we need
	 * is to find the element with both the right score and nodeect. */
	x = x.level[0].forward
	if x != nil && score == x.score && x.key == key {
		zsl.zslDeleteNode(x, update)
		return true
	}

	return false
}

/* Find the rank for an element by both score and node.
 * Returns 0 when the element cannot be found, rank otherwise.
 * Note that the rank is 1-based due to the span of zsl->header to the
 * first element. */
func (zsl *skipList[K, S]) zslGetRank(score S, key K) int64 {
	rank := uint64(0)
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}

		/* x might be equal to zsl->header, so test if node is non-NULL */
		if x.key == key {
			return int64(rank)
		}
	}
	return 0
}

/* Finds an element by its rank. The rank argument needs to be 1-based. */
func (zsl *skipList[K, S]) zslGetElementByRank(rank uint64) *skipListNode[K, S] {
	traversed := uint64(0)
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (traversed+x.level[i].span) <= rank {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}

// New creates a new ZSet and return its pointer
func NewZSet[K comparable, S Value]() *ZSet[K, S] {
	return &ZSet[K, S]{
		dict: make(map[K]*node[K, S]),
		zsl:  zslCreate[K, S](),
	}
}

// Length returns counts of elements
func (z *ZSet[K, S]) Len() int64 {
	return z.zsl.length
}

// Set: add or update
func (z *ZSet[K, S]) Set(key K, score S, data ...K) {
	v, ok := z.dict[key]

	node := &node[K, S]{score: score}
	if len(data) > 0 {
		node.data = data[0]
	}
	z.dict[key] = node

	if ok {
		// when score change
		if score != v.score {
			z.zsl.zslDelete(v.score, key)
			z.zsl.zslInsert(score, key)
		}
	} else {
		z.zsl.zslInsert(score, key)
	}
}

// IncrBy
func (z *ZSet[K, S]) IncrBy(key K, score S) *node[K, S] {
	v, ok := z.dict[key]
	if !ok {
		return nil
	}

	z.zsl.zslDelete(v.score, key)
	v.score += score
	z.zsl.zslInsert(v.score, key)

	return v
}

// Delete: delete element by key
func (z *ZSet[K, S]) Delete(key K) (ok bool) {
	v, ok := z.dict[key]
	if ok {
		z.zsl.zslDelete(v.score, key)
		delete(z.dict, key)
		return true
	}
	return false
}

// GetRank returns position,score and extra data of an element which
// found by the parameter key.
// The parameter reverse determines the rank is descent or ascend，
// true means descend and false means ascend.
func (z *ZSet[K, S]) GetRank(key K, reverse bool) (int64, *node[K, S]) {
	v, ok := z.dict[key]
	if !ok {
		return -1, nil
	}
	r := z.zsl.zslGetRank(v.score, key)
	if reverse {
		r = z.zsl.length - r
	} else {
		r--
	}
	return int64(r), v
}

// GetDataByRank returns the id,score and extra data of an element which
// found by position in the rank.
// The parameter rank is the position, reverse says if in the descend rank.
func (z *ZSet[K, S]) GetDataByRank(rank int64, reverse bool) *node[K, S] {
	if rank < 0 || rank > z.zsl.length {
		return nil
	}
	if reverse {
		rank = z.zsl.length - rank
	} else {
		rank++
	}
	n := z.zsl.zslGetElementByRank(uint64(rank))
	if n == nil {
		return nil
	}

	return z.dict[n.key]
}

// Range
func (z *ZSet[K, S]) Range(start, end int64, f func(S, K, any)) {
	z.commonRange(start, end, false, f)
}

// RevRange
func (z *ZSet[K, S]) RevRange(start, end int64, f func(S, K, any)) {
	z.commonRange(start, end, true, f)
}

func (z *ZSet[K, S]) commonRange(start, end int64, reverse bool, f func(S, K, any)) {
	l := z.zsl.length
	if start < 0 {
		start += l
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end += l
	}

	if start > end || start >= l {
		return
	}
	if end >= l {
		end = l - 1
	}
	span := (end - start) + 1

	var node *skipListNode[K, S]
	if reverse {
		node = z.zsl.tail
		if start > 0 {
			node = z.zsl.zslGetElementByRank(uint64(l - start))
		}
	} else {
		node = z.zsl.header.level[0].forward
		if start > 0 {
			node = z.zsl.zslGetElementByRank(uint64(start + 1))
		}
	}
	for span > 0 {
		span--
		k := node.key
		s := node.score
		f(s, k, z.dict[k].data)
		if reverse {
			node = node.backward
		} else {
			node = node.level[0].forward
		}
	}
}

func (z *ZSet[K, S]) Print() {
	fmt.Println(z.Len())
	for k, v := range z.dict {
		fmt.Println(k, v)
	}
}
