package intset

import (
	"fmt"
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

// exercise 6.2
func TestIntSetAddAll(t *testing.T) {
	var intSet IntSet
	elements := []int{1, 2, 100}
	intSet.AddAll(elements...)
	for i := range elements {
		if !intSet.Has(elements[i]) {
			t.Errorf("%d is not in the set", elements[i])
		}
	}

}

// exercise 6.3
func TestIntSetIntersectWith(t *testing.T) {
	var setA, setB IntSet
	setA.AddAll(1, 2, 2000, 10000)
	setB.AddAll(2, 2000, 3000)
	setA.IntersectWith(&setB)
	fmt.Println(&setA)
	if setA.Len() != 2 {
		t.Errorf("setA has %d elements", setA.Len())
	}
	if setA.Has(10000) {
		t.Errorf("setA has 10000")
	}
	if !setA.Has(2000) {
		t.Errorf("setA does not have 2000")
	}
}

func TestIntSetDifferenceWith(t *testing.T) {
	var setA, setB IntSet
	setA.AddAll(1, 2, 2000, 5000)
	setB.AddAll(2, 2000, 3000)
	setA.DifferenceWith(&setB)
	if setA.Len() != 2 {
		t.Log("setA", &setA)
		t.Errorf("setA has %d elements", setA.Len())
	}
	if !setA.Has(5000) {
		t.Errorf("setA does not have 5000")
	}
}

func TestIntSetDifferenceWith2(t *testing.T) {
	var setA, setB IntSet
	setA.AddAll(1, 2, 2000, 5000)
	setB.AddAll(2, 2000, 3000)
	setB.DifferenceWith(&setA)
	if setB.Len() != 1 {
		t.Log("setB", &setA)
		t.Errorf("setA has %d elements", setA.Len())
	}
	if !setB.Has(3000) {
		t.Errorf("setB does not have 3000")
	}
}

func TestIntSetSymmetricDifferenceWith(t *testing.T) {
	var setA, setB IntSet
	setA.AddAll(1, 2, 2000, 5000)
	setB.AddAll(2, 2000, 3000)
	setA.SymmetricDifferenceWith(&setB)
	if setA.Len() != 3 {
		t.Log("setA", &setA)
		t.Errorf("setA has %d elements", setA.Len())
	}
	if !setA.Has(1) {
		t.Errorf("setA does not have 1")
	}
	if !setA.Has(3000) {
		t.Errorf("setA does not have 1")
	}
	if !setA.Has(3000) {
		t.Errorf("setA does not have 1")
	}
}

func TestIntSetElements(t *testing.T) {
	var set IntSet
	expectedElements := []int{1, 2, 1000, 10000}
	set.AddAll(expectedElements...)
	elements := set.Elems()
	if len(elements) != set.Len() {
		t.Error("len(elements) != set.Len()")
	}
	for i := range elements {
		found := false
		for j := range expectedElements {
			if expectedElements[j] == elements[i] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%d is not among the elemnts", elements[i])
		}
	}
}
