package pokecache

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"
)

func randomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func TestCacheFunctioning(t *testing.T) {
	cache := NewCache(2 * time.Second)
	for i := 0; i < 10; i++ {
		cache.Add(fmt.Sprintf("%v index", i), randomBytes(25))
	}
	if len(cache.CacheData) != 10 {
		t.Errorf("Cache data not stored")
	}
	time.Sleep(3 * time.Second)
	if len(cache.CacheData) != 0 {
		t.Errorf("Cache data not clearning up on schedule.")
	}
}
