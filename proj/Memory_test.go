package main

import (
	"fmt"
	"runtime"
	"testing"
)

type Human struct {
	Name   string
	Age    int
	Height int
	Weight float64
}

func measureHeap() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc
}

func TestHeapAllocation(t *testing.T) {
	// Заміряємо початковий обсяг хіпу
	initialHeap := measureHeap()
	fmt.Printf("Initial Heap: %d bytes\n", initialHeap)

	// Створюємо об'єкт на стеку
	humanOnStack := Human{
		Name:   "John",
		Age:    30,
		Height: 180,
		Weight: 75.5,
	}

	// Використовуємо змінну humanOnStack, щоб уникнути помилки компіляції
	fmt.Printf("Human on Stack: %+v\n", humanOnStack)

	// Збираємо сміття, щоб точніше виміряти хіп після створення об'єкта на стеку
	runtime.GC()

	stackHeap := measureHeap()
	fmt.Printf("Heap after stack allocation: %d bytes\n", stackHeap)

	// Перевіряємо, що обсяг хіпу не змінився
	if stackHeap != initialHeap {
		t.Errorf("Heap allocation changed after stack allocation. Initial: %d, After: %d", initialHeap, stackHeap)
	}

	// Створюємо об'єкт на хіпі
	humanOnHeap := &Human{
		Name:   "Jane",
		Age:    25,
		Height: 165,
		Weight: 60.0,
	}

	// Використовуємо змінну humanOnHeap, щоб уникнути помилки компіляції
	fmt.Printf("Human on Heap: %+v\n", *humanOnHeap)

	// Збираємо сміття, щоб точніше виміряти хіп після створення об'єкта на хіпі
	runtime.GC()

	heapHeap := measureHeap()
	fmt.Printf("Heap after heap allocation: %d bytes\n", heapHeap)

	// Перевіряємо, що обсяг хіпу збільшився
	if heapHeap <= stackHeap {
		t.Errorf("Heap allocation did not increase after heap allocation. After stack allocation: %d, After heap allocation: %d", stackHeap, heapHeap)
	}
}
