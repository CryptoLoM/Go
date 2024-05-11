package main

import (
    
    "runtime"
    "testing"

	
)
func TestGarbageCollector(t *testing.T) {
    // Нагенеруємо сміття
    for i := 0; i < 10000; i++ {
        _ = make([]byte, 1024)
    }

    // Запам'ятовуємо початковий розмір хіпу
    var stats runtime.MemStats
    runtime.ReadMemStats(&stats)
    heapSizeBefore := stats.HeapAlloc

    // Викликаємо garbage collector
    runtime.GC()

    // Знову зчитуємо розмір хіпу після збирання сміття
    runtime.ReadMemStats(&stats)
    heapSizeAfter := stats.HeapAlloc

    // Перевірка, чи розмір хіпу зменшився
    if heapSizeAfter >= heapSizeBefore {
        t.Errorf("Помилка: очікувалось, що розмір хіпу зменшиться після виклику garbage collector")
    }
}