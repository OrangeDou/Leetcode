package main

import "sync"

type Shard struct {
	sync.RWMutex
	data map[string]interface{}
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
			data: make(map[string]interface{}),
		}
	}
	return &ShardMap{
		shards:   Shards,
		numShard: numsShard,
	}
}

//
