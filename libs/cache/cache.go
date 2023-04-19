package cache

import (
	"container/list"
	"sync"
	"time"
)

type CacheEntry[T any] struct {
	key      string
	deadline time.Time
	value    *T
}

type Cache[T any] struct {
	buckets    []Bucket[T]
	bucketsCnt int32
	ttl        time.Duration
}

// вместимость одного бакета
const (
	bucketCapacity = int32(16)
	bucketsCount   = int32(16)
)

type Bucket[T any] struct {
	// locks bucket when reading
	mx *sync.RWMutex
	// priority queue for LRU
	pq *list.List
	// storage with cache entries
	storage map[int32]*list.List
}

func (bucket *Bucket[T]) clearOutdated() {
	now := time.Now()

	bucket.mx.Lock()
	defer bucket.mx.Unlock()

	// run over bucket entries
	for _, val := range bucket.storage {
		// if entry is empty, skip it
		if val == nil {
			continue
		}

		// iterate over entry elements
		for curr := val.Front(); curr != nil; curr = curr.Next() {
			currentValue := curr.Value.(CacheEntry[T])

			// if element is outdated, remove it from the list
			if now.After(currentValue.deadline) {
				val.Remove(curr)
			}
		}
	}
}

func (bucket *Bucket[T]) findElementInStorage(key string, hash int32) *list.Element {
	list := bucket.storage[hash]
	if list == nil {
		return nil
	}

	for element := list.Front(); element != nil; element = element.Next() {
		val := element.Value.(CacheEntry[T])

		if val.key == key {
			// update position on element read
			bucket.pq.MoveToFront(element)
			return element
		}
	}

	return nil
}

func (bucket *Bucket[T]) getElement(key string, hash int32) *list.Element {
	bucket.mx.RLock()
	defer bucket.mx.RUnlock()

	return bucket.findElementInStorage(key, hash)
}

func (bucket *Bucket[T]) addElement(key string, value CacheEntry[T]) {
	hash := hashFAQ6(key)

	bucket.mx.Lock()
	defer bucket.mx.Unlock()

	// bucket overflow, clear least used element
	if bucket.pq.Len() >= int(bucketCapacity) {
		lastEl := bucket.pq.Back()
		bucket.pq.Remove(lastEl)

		val := lastEl.Value.(CacheEntry[T])
		valHash := hashFAQ6(val.key)

		elInStorage := bucket.findElementInStorage(val.key, valHash)
		bucket.storage[valHash].Remove(elInStorage)
	}

	bucketList := bucket.storage[hash]
	if bucketList == nil {
		bucketList = &list.List{}
		bucket.storage[hash] = bucketList
	}

	// new cache is most recent, push it to the front
	bucketList.PushFront(value)
	bucket.pq.PushFront(value)
}

// TODO: здесь можно переделать на consistent hashing, пока базовый алгоритм
// determine bucket by first letter
func (cache *Cache[T]) getBucketIdx(key string) int32 {
	return int32(key[0] % byte(bucketsCount))
}

func (cache *Cache[T]) AddToCache(key string, value T) {
	bucketIdx := cache.getBucketIdx(key)

	cacheEntry := CacheEntry[T]{
		key:      key,
		value:    &value,
		deadline: time.Now().Add(cache.ttl),
	}

	cache.buckets[bucketIdx].addElement(key, cacheEntry)
}

func (cache *Cache[T]) GetFromCache(key string) (*T, bool) {
	bucket := cache.buckets[cache.getBucketIdx(key)]
	hash := hashFAQ6(key)

	el := bucket.getElement(key, hash)
	if el == nil {
		return nil, false
	}

	val, ok := el.Value.(CacheEntry[T])

	if !ok {
		return nil, false
	}
	return val.value, true
}

func (cache *Cache[T]) executeTTL() {
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)

			for _, bucket := range cache.buckets {
				bucket.clearOutdated()
			}
		}
	}()
}

func NewCache[T any](ttl time.Duration) *Cache[T] {
	buckets := make([]Bucket[T], bucketsCount)

	for idx := range buckets {
		buckets[idx] = Bucket[T]{
			mx:      &sync.RWMutex{},
			pq:      &list.List{},
			storage: make(map[int32]*list.List),
		}
	}

	cache := &Cache[T]{
		buckets:    buckets,
		bucketsCnt: bucketsCount,
		ttl:        ttl,
	}

	cache.executeTTL()

	return cache
}
