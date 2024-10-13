package model

import (
	"testing"
	"time"
)

func TestLRUCache_SetAndGet(t *testing.T) {
	cache := NewLRUCache(2)

	// Test setting a value
	cache.Set("key1", "value1", 5*time.Second)
	value, ok := cache.Get("key1")
	if !ok || value.(*Scache).Value != "value1" {
		t.Fatalf("expected to get 'value1', got %v", value)
	}

	// Test overwriting a value
	cache.Set("key1", "value2", 5*time.Second)
	value, ok = cache.Get("key1")
	if !ok || value.(*Scache).Value != "value2" {
		t.Fatalf("expected to get 'value2', got %v", value)
	}

	// Test eviction when capacity is exceeded
	cache.Set("key2", "value2", 5*time.Second)
	cache.Set("key3", "value3", 5*time.Second) // This should evict key1

	_, ok = cache.Get("key1")
	if ok {
		t.Fatalf("expected key1 to be evicted")
	}

	value, ok = cache.Get("key2")
	if !ok || value.(*Scache).Value != "value2" {
		t.Fatalf("expected to get 'value2', got %v", value)
	}
}

func TestLRUCache_Expiration(t *testing.T) {
	cache := NewLRUCache(2)

	// Test setting a value with expiration
	cache.Set("key1", "value1", 1*time.Second)

	time.Sleep(2 * time.Second) // Wait for the key to expire

	_, ok := cache.Get("key1")
	if ok {
		t.Fatalf("expected key1 to be expired")
	}
}

func TestLRUCache_MultipleEvictions(t *testing.T) {
	cache := NewLRUCache(2)

	// Set multiple values
	cache.Set("key1", "value1", 5*time.Second)
	cache.Set("key2", "value2", 5*time.Second)
	cache.Set("key3", "value3", 5*time.Second) // Evicts key1

	_, ok := cache.Get("key1")
	if ok {
		t.Fatalf("expected key1 to be evicted")
	}

	cache.Set("key4", "value4", 5*time.Second) // Evicts key2

	_, ok = cache.Get("key2")
	if ok {
		t.Fatalf("expected key2 to be evicted")
	}
}
