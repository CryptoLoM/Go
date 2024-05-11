package main

import (
	"runtime"
	"testing"
)

func TestGarbageCollector(t *testing.T) {
	// Generate garbage
	var garbage [][]byte
	for i := 0; i < 10000; i++ {
		garbage = append(garbage, make([]byte, 1024*1024)) 
	}

	
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.GC() 
	runtime.ReadMemStats(&memStatsBefore)
	heapBefore := memStatsBefore.HeapAlloc

	// Generate more garbage
	for i := 0; i < 10000; i++ {
		garbage = append(garbage, make([]byte, 1024*1024)) 
	}

	runtime.GC()

	runtime.ReadMemStats(&memStatsAfter)
	heapAfter := memStatsAfter.HeapAlloc

	if heapAfter >= heapBefore {
		t.Errorf("Expected memory to be freed after garbage collection, but the actual heap size after collection: %d, heap size before collection: %d", heapAfter, heapBefore)
	}
}
