package intset

import (
	"testing"
)

//exercise 6.1
func TestLen(t *testing.T) {
	var intSet IntSet
	intSet.Add(1)
	intSet.Add(120)
	intSet.Add(5000)
	numberOfElements := intSet.Len()
	if numberOfElements != 3 {
		t.Errorf(`intSet.Len() == %d`, numberOfElements)
	}
}
func TestIntSetRemove(t *testing.T) {
	var intSet IntSet
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(1000)
	intSet.Remove(1)
	intSet.Remove(2)
	if intSet.Has(2) {
		t.Error("intSet.Has(2) == true ")
	}
	if !intSet.Has(1000) {
		t.Error("intSet.Has(1000) == false")
	}
	if intSet.Has(1) {
		t.Error("intSet.Has(1) == true")
	}
}

func TestIntSetClear(t *testing.T) {
	var intSet IntSet
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(1000)
	intSet.Clear()
	if intSet.Len() != 0 {
		t.Error("intSet.Len() != 0")
	}
}

func TestIntSetCopy(t *testing.T) {
	var intSet IntSet
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(1000)
	copySet := intSet.Copy()
	intSet.Clear()
	if !copySet.Has(1000) {
		t.Error()
	}
	if copySet.Len() != 3 {
		t.Error()
	}

}
