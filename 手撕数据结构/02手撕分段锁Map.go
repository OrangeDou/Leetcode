package main

import (
	"fmt"
	"sync"
)

type Shard struct {
	sync.RWMutex
	data map[string][]interface{}
}

type ShardMap struct {
	shards   []*Shard
	numShard int
}

// Newshardmap
func Newshardmap(numsShard int) *ShardMap {
	Shards := make([]*Shard, numsShard)
	for i := 0; i < numsShard; i++ {
		Shards[i] = &Shard{
			data: make(map[string][]interface{}),
		}
	}
	return &ShardMap{
		shards:   Shards,
		numShard: numsShard,
	}
}

// get
func (s *ShardMap) Get(key string) ([]interface{}, bool) {
	index := s.getShardIndex(key)
	shard := s.shards[index]
	shard.RLock()
	defer shard.RUnlock()
	value, exists := shard.data[key]
	return value, exists

}

// set
func (s *ShardMap) Set(key string, value interface{}) {
	index := s.getShardIndex(key)
	Shard := s.shards[index]
	Shard.Lock()
	defer Shard.Unlock()
	Shard.data[key] = append(Shard.data[key], value)
}

func (s *ShardMap) getShardIndex(key string) int {
	return int(key[0]) % s.numShard
}

func main() {
	shardedMap := Newshardmap(10)

	// Set some values
	shardedMap.Set("key1", "value1")
	shardedMap.Set("key2", "value2")

	// Get values
	if value, exists := shardedMap.Get("key1"); exists {
		fmt.Println("key1:", value)
	}

	// Delete a key
	shardedMap.Set("key2", nil)

	// Try to get the deleted key
	if value, exists := shardedMap.Get("key2"); exists {
		fmt.Println("key2:", value)
	} else {
		fmt.Println("key2 does not exist")
	}
}
