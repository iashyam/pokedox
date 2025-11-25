package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	// Test Get
	if val, ok := cache.Get(key); !ok || string(val) != "testValue" {
		t.Errorf("Expected value %s, got %s", "testValue", string(val))
	}
}

func TestReap(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)
	cache.Reap(key)

	// Test Get after Reap
	if _, ok := cache.Get(key); ok {
		t.Errorf("Expected key %s to be reaped, but it still exists", key)
	}
}

func TestReapAll(t *testing.T) {
	cache := NewCache()
	key1 := "key1"
	key2 := "key2"
	value := []byte("value")

	cache.Add(key1, value)
	time.Sleep(2 * time.Second) // Simulate delay
	cache.Add(key2, value)

	// Reap all entries older than 1 second
	go cache.ReapAll(1 * time.Second)
	time.Sleep(3 * time.Second) // Allow ReapAll to run

	if _, ok := cache.Get(key1); ok {
		t.Errorf("Expected key %s to be reaped, but it still exists", key1)
	}

	if _, ok := cache.Get(key2); !ok {
		t.Errorf("Expected key %s to exist, but it was reaped", key2)
	}
}

