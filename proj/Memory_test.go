package main

import (
    "fmt"
    "testing"
)
//створюємо клас
type Human struct {
    Name   string
    Age    int
    Height int
    Weight float64
}

func TestStackToHeap(t *testing.T) {
    // Створення якогось тіпочка на стеку
    human := Human{"John", 30, 182, 80.5}
    
    // Отримання адреси пам'яті тіпочка до створення
    stackBoundaryBefore := fmt.Sprintf("%p", &human)

    // Отримання адреси пам'яті тіпочка після створення
    stackBoundaryAfter := fmt.Sprintf("%p", &human)

    // Перевірка, чи різниця на хіпі до і після створення буде нульовою
    if stackBoundaryBefore != stackBoundaryAfter {
        t.Errorf("Помилка: різниця на хіпі до і після створення тіпочка на стеці не нульова")
    }
}
func TestHeapToStack(t *testing.T) {
    // Створення тіпочка на купі (використовуючи вказівник)
    humanPtr := &Human{"Jane", 25, 167, 53}

    // Отримання адреси пам'яті тіпочка на купі
    heapBoundaryBefore := fmt.Sprintf("%p", humanPtr)

    // Створення тіпочка на стеку
    human := Human{"John", 30, 182, 80.5}

    // Отримання адреси пам'яті тіпочка на стеку
    stackBoundaryBefore := fmt.Sprintf("%p", &human)

    // Перевірка, чи різниця між адресами пам'яті до і після створення не нульова
    if stackBoundaryBefore == heapBoundaryBefore {
        t.Errorf("Помилка: різниця на хіпі до і після створення тіпочка на хіпі нульова")
    }
}
