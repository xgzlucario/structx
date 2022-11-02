package structx

import (
	"math/rand"
)

const zSkiplistMaxlevel = 32

type (
	skipListLevel struct {
		forward *skipListNode
		span    uint64
	}
	skipListNode struct {
		objID    int64
		score    float64
		backward *skipListNode
		level    []*skipListLevel
	}
	obj struct {
		key   int64
		data  any
		score float64
	}
	skipList struct {
		header *skipListNode
		tail   *skipListNode
		length int64
		level  int16
	}

	ZSet struct {
		dict map[int64]*obj
		zsl  *skipList
	}
)

func zslCreateNode(level int16, score float64, id int64) *skipListNode {
	n := &skipListNode{
		score: score,
		objID: id,
		level: make([]*skipListLevel, level),
	}
	for i := range n.level {
		n.level[i] = new(skipListLevel)
	}
	return n
}

func zslCreate() *skipList {
	return &skipList{
		level:  1,
		header: zslCreateNode(zSkiplistMaxlevel, 0, 0),
	}
}

const zSkiplistP = 0.25 /* Skiplist P = 1/4 */

/* Returns a random level for the new skiplist node we are going to create.
 * The return value of this function is between 1 and _ZSKIPLIST_MAXLEVEL
 * (both inclusive), with a powerlaw-alike distribution where higher
 * levels are less likely to be returned. */
func randomLevel() int16 {
	level := int16(1)
	for float32(rand.Int31()&0xFFFF) < (zSkiplistP * 0xFFFF) {
		level++
	}
	if level < zSkiplistMaxlevel {
		return level
	}
	return zSkiplistMaxlevel
}

/* zslInsert a new node in the skiplist. Assumes the element does not already
 * exist (up to the caller to enforce that). The skiplist takes ownership
 * of the passed SDS string 'obj'. */
func (zsl *skipList) zslInsert(score float64, id int64) *skipListNode {
	update := make([]*skipListNode, zSkiplistMaxlevel)
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
					(x.level[i].forward.score == score && x.level[i].forward.objID < id)) {
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
	x = zslCreateNode(level, score, id)
	for i := int16(0); i < level; i++ {
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
func (zsl *skipList) zslDeleteNode(x *skipListNode, update []*skipListNode) {
	for i := int16(0); i < zsl.level; i++ {
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
 * referenced SDS string at node->obj). */
func (zsl *skipList) zslDelete(score float64, id int64) int {
	update := make([]*skipListNode, zSkiplistMaxlevel)
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					x.level[i].forward.objID < id)) {
			x = x.level[i].forward
		}
		update[i] = x
	}
	/* We may have multiple elements with the same score, what we need
	 * is to find the element with both the right score and object. */
	x = x.level[0].forward
	if x != nil && score == x.score && x.objID == id {
		zsl.zslDeleteNode(x, update)
		return 1
	}
	return 0 /* not found */
}

/* Find the rank for an element by both score and obj.
 * Returns 0 when the element cannot be found, rank otherwise.
 * Note that the rank is 1-based due to the span of zsl->header to the
 * first element. */
func (zsl *skipList) zslGetRank(score float64, key int64) int64 {
	rank := uint64(0)
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					x.level[i].forward.objID <= key)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}

		/* x might be equal to zsl->header, so test if obj is non-NULL */
		if x.objID == key {
			return int64(rank)
		}
	}
	return 0
}

/* Finds an element by its rank. The rank argument needs to be 1-based. */
func (zsl *skipList) zslGetElementByRank(rank uint64) *skipListNode {
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
func New() *ZSet {
	return &ZSet{
		dict: make(map[int64]*obj),
		zsl:  zslCreate(),
	}
}

// Length returns counts of elements
func (z *ZSet) Len() int64 {
	return z.zsl.length
}

// Set is used to add or update an element
func (z *ZSet) Set(score float64, key int64, dat any) {
	v, ok := z.dict[key]
	z.dict[key] = &obj{data: dat, key: key, score: score}
	if ok {
		/* Remove and re-insert when score changes. */
		if score != v.score {
			z.zsl.zslDelete(v.score, key)
			z.zsl.zslInsert(score, key)
		}
	} else {
		z.zsl.zslInsert(score, key)
	}
}

// IncrBy ..
func (z *ZSet) IncrBy(score float64, key int64) (float64, any) {
	v, ok := z.dict[key]
	if !ok {
		// use negative infinity ?
		return 0, nil
	}
	if score != 0 {
		z.zsl.zslDelete(v.score, key)
		v.score += score
		z.zsl.zslInsert(v.score, key)
	}
	return v.score, v.data
}

// Delete removes an element from the ZSet
// by its key.
func (z *ZSet) Delete(key int64) (ok bool) {
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
func (z *ZSet) GetRank(key int64, reverse bool) (rank int64, score float64, data any) {
	v, ok := z.dict[key]
	if !ok {
		return -1, 0, nil
	}
	r := z.zsl.zslGetRank(v.score, key)
	if reverse {
		r = z.zsl.length - r
	} else {
		r--
	}
	return int64(r), v.score, v.data

}

// GetData returns data stored in the map by its key
func (z *ZSet) GetData(key int64) (data any, ok bool) {
	o, ok := z.dict[key]
	if !ok {
		return nil, false
	}
	return o.data, true
}

// GetScore implements ZScore
func (z *ZSet) GetScore(key int64) (score float64, ok bool) {
	o, ok := z.dict[key]
	if !ok {
		return 0, false
	}
	return o.score, true
}

// GetDataByRank returns the id,score and extra data of an element which
// found by position in the rank.
// The parameter rank is the position, reverse says if in the descend rank.
func (z *ZSet) GetDataByRank(rank int64, reverse bool) (key int64, score float64, data any) {
	if rank < 0 || rank > z.zsl.length {
		return 0, 0, nil
	}
	if reverse {
		rank = z.zsl.length - rank
	} else {
		rank++
	}
	n := z.zsl.zslGetElementByRank(uint64(rank))
	if n == nil {
		return 0, 0, nil
	}
	dat, ok := z.dict[n.objID]
	if !ok {
		return 0, 0, nil
	}
	return dat.key, dat.score, dat.data
}

// Range implements ZRANGE
func (z *ZSet) Range(start, end int64, f func(float64, int64, any)) {
	z.commonRange(start, end, false, f)
}

// RevRange implements ZREVRANGE
func (z *ZSet) RevRange(start, end int64, f func(float64, int64, any)) {
	z.commonRange(start, end, true, f)
}

func (z *ZSet) commonRange(start, end int64, reverse bool, f func(float64, int64, any)) {
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

	var node *skipListNode
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
		k := node.objID
		s := node.score
		f(s, k, z.dict[k].data)
		if reverse {
			node = node.backward
		} else {
			node = node.level[0].forward
		}
	}
}
