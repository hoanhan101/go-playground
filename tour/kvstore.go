//
// KVStore
//
package main

import (
    "fmt"
    "sync"
    "time"
)

type KVStore struct {
    kv  map[string]int
    mu  sync.Mutex
}

func (s *KVStore) inc(key string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.kv[key]++
}

func (s *KVStore) put(key string, value int) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.kv[key] = value
}

func (s *KVStore) get(key string) int {
    s.mu.Lock()
    defer s.mu.Unlock()

    return s.kv[key]
}

func (s *KVStore) del(key string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    delete(s.kv, key)
}

func main() {
    kvs := KVStore{kv: make(map[string]int)}

    kvs.put("test", 1)

    go func() {
        for i := 0; i < 10; i++ {
            kvs.inc("test")
        }
    }()

    go func() {
        for i := 0; i < 10; i++ {
            kvs.inc("test")
        }
    }()

    time.Sleep(time.Second)
    fmt.Println(kvs.get("test"))
    
    kvs.del("test")
    fmt.Println(kvs.get("test"))
}
