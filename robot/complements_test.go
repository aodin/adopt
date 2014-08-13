package robot

import "testing"

func TestComplements(t *testing.T) {
	// Test equal
	a := []int64{1, 2, 3}
	b := []int64{3, 2, 1}
	mFromA, mFromB := Complements(a, b)
	if len(mFromA) != 0 || len(mFromB) != 0 {
		t.Errorf("Equal arrays should have no complements: %v, %v", mFromA, mFromB)
	}

	// Test one difference
	a = []int64{1, 2, 4}
	b = []int64{3, 2, 1}
	mFromA, mFromB = Complements(a, b)
	if len(mFromA) != 1 || len(mFromB) != 1 {
		t.Fatalf("There should be complements: %v, %v", mFromA, mFromB)
		if mFromA[0] != 3 {
			t.Errorf("3 should be missing from A")
		}
		if mFromB[0] != 4 {
			t.Errorf("4 should be missing from B")
		}
	}

	// Test major differences
	a = []int64{5, 6, 7}
	b = []int64{}
	mFromA, mFromB = Complements(a, b)
	if len(mFromA) != 0 || len(mFromB) != 3 {
		t.Fatalf("There should be complements for b: %v, %v", mFromB)
		if mFromB[0] != 5 {
			t.Errorf("5 should be missing from B")
		}
		if mFromB[1] != 6 {
			t.Errorf("5 should be missing from B")
		}
		if mFromB[2] != 7 {
			t.Errorf("5 should be missing from B")
		}
	}

	a = []int64{}
	b = []int64{5, 6, 7}
	mFromA, mFromB = Complements(a, b)
	if len(mFromA) != 3 || len(mFromB) != 0 {
		t.Fatalf("There should be complements for a: %v, %v", mFromB)
		if mFromA[0] != 5 {
			t.Errorf("5 should be missing from A")
		}
		if mFromA[1] != 6 {
			t.Errorf("5 should be missing from A")
		}
		if mFromA[2] != 7 {
			t.Errorf("5 should be missing from A")
		}
	}
}
