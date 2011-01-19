package main

import "sync"

var (
	uniqueIdMutex sync.Mutex
	uniqueId uint64
)

func GenerateUniqueID() uint64 {
	uniqueIdMutex.Lock()
	defer uniqueIdMutex.Unlock()
	uniqueId++
	
	return uniqueId
}
