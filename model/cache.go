package model 


import (
	"sync"
	"time"
)

type Scache struct {
	Key string
	Value interface{}
	Expiration int64
	PrevCache  *Scache
	NextCache *Scache

}

type LRUCache struct{
	MaxSize int
	Cache map[string]*Scache
	head *Scache
	tail *Scache
	mu sync.RWMutex
}

func  NewLRUCache(capacity int)  *LRUCache{
	head := &Scache{}
	tail := &Scache{}
	head.NextCache = tail
	tail.NextCache = head

	return &LRUCache{
		MaxSize: capacity,
		Cache:    make(map[string]*Scache),
		head:     head,
		tail:     tail,
	}
}
//Head <-> A <-> B <-> C <-> Tail

func (l *LRUCache) MoveToFront(cacheItem *Scache) {
	cacheItem.NextCache.PrevCache = cacheItem.NextCache
	cacheItem.NextCache.PrevCache = cacheItem.PrevCache

	cacheItem.NextCache = l.head.NextCache
	cacheItem.PrevCache = l.head
	
	l.head.NextCache.PrevCache = cacheItem
	l.head.NextCache = cacheItem
}

func (l * LRUCache) Set(key string, value interface{}, duration time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if item , exist := l.Cache[key]; exist{
		item.Value = value
		item.Expiration = time.Now().Add(duration).UnixNano()
		l.MoveToFront(item)
		return 
	}

	newItem := &Scache{
		Key:        key,
		Value:      value,
		Expiration: time.Now().Add(duration).UnixNano(),
	}

	l.Cache[key] = newItem
	l.MoveToFront(newItem)
	if len(l.Cache) > l.MaxSize {
		l.removeTail()
	}
}


func (l *LRUCache) removeTail() {
	if l.tail.PrevCache == l.head {
		return // No items to remove
	}

	lruItem := l.tail.PrevCache
	l.tail.PrevCache = lruItem.PrevCache
	lruItem.PrevCache.NextCache = l.tail
	delete(l.Cache, lruItem.Key) // Adjusting this to delete by Key
}
