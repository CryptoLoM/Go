package main

import (
    "testing"
)
// Тест для демонстрації роботи смітника
func TestGarbageCollection(t *testing.T) {
    for i := 0; i < 1000000; i++ {
        
        _ = new(ReferenceType)
    }
   
}