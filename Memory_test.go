package main

import (
    "fmt"
    "testing"
)
func TestMemoryAllocation(t *testing.T) {
    // Створення об'єктів на стеку
    var valOnStack ValueType
    fmt.Println("Value type on stack:", valOnStack)

    // Створення об'єктів на хіпі
    refOnHeap := new(ReferenceType)
    fmt.Println("Reference type on heap:", refOnHeap)
}

