package server

import (
	"sync"
	"fmt"
)

type KeyValueStore struct{
	store map[string]string
	mutex sync.RWMutex
}

type ReturnValue struct{
	Ok bool
	Value string
}

func (r *ReturnValue) String() string{
	return fmt.Sprintf("Exist: %v; Value:%v", r.Ok, r.Value)
}

type Pair struct{
	Key string
	Value string
}

func NewStore() *KeyValueStore{
	return &KeyValueStore{
		store: make(map[string]string)}
}

func (k *KeyValueStore) Get(key string, result *ReturnValue) error{
	k.mutex.RLock()
	v, ok := k.store[key]
	k.mutex.RUnlock()
	result.Ok = ok
	result.Value = v
	return nil 
}

func (k *KeyValueStore) Set(entry Pair, ok *bool) error{
	k.mutex.Lock()
	k.store[entry.Key] = entry.Value
	k.mutex.Unlock()
	*ok = true
	return nil
} 