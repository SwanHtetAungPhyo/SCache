package model

import (
	"encoding/json"
	"os"
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
	if cacheItem == nil || cacheItem.PrevCache == nil || cacheItem.NextCache == nil {
		return // Prevent nil dereference
	}


	cacheItem.PrevCache.NextCache = cacheItem.NextCache
	cacheItem.NextCache.PrevCache = cacheItem.PrevCache

	cacheItem.NextCache = l.head.NextCache
	cacheItem.PrevCache = l.head
	l.head.NextCache.PrevCache = cacheItem
	l.head.NextCache = cacheItem
}

func (l *LRUCache) Set(key string, value interface{}, duration time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if item, exist := l.Cache[key]; exist {
		item.Value = value
		item.Expiration = time.Now().Add(duration).UnixNano()
		l.MoveToFront(item)
		return
	}

	newItem := &Scache{
		Key:        key,
		Value:      value,
		Expiration: time.Now().Add(duration).UnixNano(),
		PrevCache:  nil, 
		NextCache:  nil, 
	}


	l.Cache[key] = newItem
	l.MoveToFront(newItem)

	if len(l.Cache) > l.MaxSize {
		l.removeTail()
	}
}



func (l *LRUCache) removeTail() {
	if l.tail.PrevCache == l.head {
		return
	}

	lruItem := l.tail.PrevCache
	l.tail.PrevCache = lruItem.PrevCache
	lruItem.PrevCache.NextCache = l.tail
	delete(l.Cache, lruItem.Key) 
}

func (l *LRUCache) Get(key string) (interface{}, bool){
	l.mu.Lock()
	defer l.mu.Unlock()

	 item, exist := l.Cache[key]; 
	 if !exist{
		return item, false
	}

	if item.Expiration < time.Now().UnixNano(){
		l.Evict(key)
		return nil , false
	}
	return item, true
}

func (l * LRUCache) Evict(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	item, exist := l.Cache[key]
	if !exist{
		return
	}
	item.PrevCache.NextCache = item.NextCache
	item.NextCache.PrevCache = item.PrevCache

	delete(l.Cache, key)
}


func (l *LRUCache) InternalClearance(){
	l.mu.Lock()
	defer l.mu.Unlock()

	for key, item := range l.Cache{
		if item.Expiration < time.Now().UnixNano(){
			l.Evict(key)
		}
	}
}

func (l *LRUCache) SnapShoter(filepath string) error{
	l.mu.Lock()
	defer l.mu.Unlock()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(l.Cache)
}

