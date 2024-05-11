package main

import (
    "testing"
)

type ValueType struct {
    val int
}

type ReferenceType struct {
    val *int
}

func TestValueVsReference(t *testing.T) {
    // Value types передаються за значенням
    a := ValueType{10}
    modifyValue(a)
    if a.val != 10 {
        t.Error("Value types were modified")
    }

    // Reference types передаються за посиланням
    b := ReferenceType{val: new(int)}
    *b.val = 10
    modifyReference(&b)
    if *b.val != 20 {
        t.Error("Reference types were not modified")
    }
}

func modifyValue(v ValueType) {
    v.val = 20
}

func modifyReference(r *ReferenceType) {
    *r.val = 20
}




