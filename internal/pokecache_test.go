package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(10)
	cache.Add("key", []byte("value"))
	val, ok := cache.Get("key")
	if !ok {
		t.Fatalf("key not found")
	}
	if string(val) != "value" {
		t.Fatalf("value mismatch")
	}

}

func TestReapLoop(t *testing.T) {
	cache := NewCache(10)
	cache.Add("key", []byte("value"))
	time.Sleep(20 * time.Millisecond)
	_, ok := cache.Get("key")
	if ok {
		t.Fatalf("key not reaped")
	}
}
