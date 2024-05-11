package main

import (
    "runtime"
    "testing"
)

func TestGarbageCollector(t *testing.T) {
    // Створення об'єктів для генерації сміття
    var garbage [][]byte
    for i := 0; i < 10000; i++ {
        // Створення сміття (повторювані байтові зрізи)
        garbage = append(garbage, make([]byte, 1024*1024))
    }

    // Вимірюємо розмір хіпу до виклику збірки сміття
    memBefore := runtime.MemStats{}
    runtime.ReadMemStats(&memBefore)
    heapBefore := memBefore.HeapAlloc

    // Викликаємо збірку сміття
    runtime.GC()

    // Вимірюємо розмір хіпу після виклику збірки сміття
    memAfter := runtime.MemStats{}
    runtime.ReadMemStats(&memAfter)
    heapAfter := memAfter.HeapAlloc

    // Перевірка, що після збірки сміття пам'ять вивільнилась (розмір хіпу зменшився)
    if heapAfter >= heapBefore {
        t.Errorf("Помилка: очікувалось, що розмір хіпу після збірки сміття буде меншим, але фактичний розмір хіпу після збірки: %d, розмір хіпу до збірки: %d", heapAfter, heapBefore)
    }
}
