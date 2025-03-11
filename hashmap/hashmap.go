package hashmap

import "github.com/gaarutyunov/galgo/linkedlist"

type pair[K comparable, V any] struct {
	key   K
	value V
}

type HashFn[K comparable] func(K) uint64

type HashMap[K comparable, V any] struct {
	buckets []*linkedlist.Node[pair[K, V]]
	hashFn  HashFn[K]
}

func New[K comparable, V any](size int, hashFn HashFn[K]) *HashMap[K, V] {
	return &HashMap[K, V]{
		buckets: make([]*linkedlist.Node[pair[K, V]], size),
		hashFn:  hashFn,
	}
}

func (h *HashMap[K, V]) Len() (n int) {
	for _, bucket := range h.buckets {
		if bucket == nil {
			continue
		}

		n += bucket.Len()
	}
	return
}

func (h *HashMap[K, V]) Set(key K, value V) {
	bucketIdx := int(h.hashFn(key)) % len(h.buckets)
	if h.buckets[bucketIdx] == nil {
		h.buckets[bucketIdx] = linkedlist.New(pair[K, V]{key, value})
	} else {
		bucket := h.buckets[bucketIdx]
		prev := bucket
		k := bucket.Value().key
		for k != key {
			prev = bucket
			bucket = bucket.Next()
			if bucket == nil {
				prev.Append(pair[K, V]{key, value})
				return
			}
			k = bucket.Value().key
		}
		if bucket == prev {
			node := linkedlist.New(pair[K, V]{key, value})
			node.AppendNode(bucket.Next())
			h.buckets[bucketIdx] = node
		} else {
			prev.ReplaceNext(pair[K, V]{key, value})
		}
	}
}

func (h *HashMap[K, V]) Get(key K) (value V, ok bool) {
	bucketIdx := int(h.hashFn(key)) % len(h.buckets)
	if h.buckets[bucketIdx] == nil {
		return
	}

	bucket := h.buckets[bucketIdx]
	k := bucket.Value().key
	for k != key {
		bucket = bucket.Next()
		if bucket == nil {
			return
		}
		k = bucket.Value().key
	}

	return bucket.Value().value, true
}

func (h *HashMap[K, V]) Delete(key K) {
	bucketIdx := int(h.hashFn(key)) % len(h.buckets)
	if h.buckets[bucketIdx] == nil {
		return
	}

	bucket := h.buckets[bucketIdx]
	prev := bucket
	k := bucket.Value().key
	for k != key {
		prev = bucket
		bucket = bucket.Next()
		if bucket == nil {
			return
		}
		k = bucket.Value().key
	}
	if prev == bucket {
		if !bucket.HasNext() {
			h.buckets[bucketIdx] = nil
		} else {
			h.buckets[bucketIdx] = bucket.Next()
		}
	} else {
		prev.SkipNext()
	}
}
