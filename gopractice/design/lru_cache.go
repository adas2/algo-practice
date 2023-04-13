package design

import "container/list"

type LRUCache struct {
	// linked list of all items caches
	values *list.List
	// key: ptr to value in list
	indexMap map[int]*list.Element
	maxCap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		values:   list.New(),
		indexMap: make(map[int]*list.Element),
		maxCap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// check if the key is present in map
	if _, exists := this.indexMap[key]; exists {
		// get the value
		elem := this.indexMap[key]
		val := elem.Value

		// move this to front of list to make it most recently used
		this.values.MoveToFront(elem)

		// update key map
		this.indexMap[key] = this.values.Front()

		return val.(entry).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// check if the key is present in map
	if _, exists := this.indexMap[key]; exists {
		elem := this.indexMap[key]
		this.values.Remove(elem)

		// add this to front of list to make it most recently used
		this.values.PushFront(entry{key, value})

		// update key map
		this.indexMap[key] = this.values.Front()

	} else {

		// key does not exists in map; and there is capacity
		if this.values.Len() < this.maxCap {

			this.values.PushFront(entry{key, value})

			this.indexMap[key] = this.values.Front()

		} else {
			// no capacity left in cache
			// remove last element, which is also least recently used elem
			// important: remove from key map
			lastKey := this.values.Back().Value.(entry).key
			delete(this.indexMap, lastKey)

			this.values.Remove(this.values.Back())

			this.values.PushFront(entry{key, value})

			this.indexMap[key] = this.values.Front()
		}
	}
}

type entry struct {
	key, value int
}
