package robot

import "sort"

type Int64Slice []int64

func (a Int64Slice) Len() int {
	return len(a)
}

func (a Int64Slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Int64Slice) Less(i, j int) bool {
	return a[i] < a[j]
}

// Int64Complements determine the complements of both int64 arrays.
func Int64Complements(a, b []int64) (missingFromA, missingFromB []int64) {
	sort.Sort(Int64Slice(a))
	sort.Sort(Int64Slice(b))
	// TODO These loops can likely be combined
	{
		var j int
		for i := 0; i < len(a) && j < len(b); {
			if a[i] < b[j] {
				missingFromB = append(missingFromB, a[i])
				i += 1
			} else if a[i] > b[j] {
				j += 1
			} else {
				i += 1
				j += 1
			}
		}
		if j < len(b) {
			missingFromA = append(missingFromA, b[j:]...)
		}
	}
	{
		var j int
		for i := 0; i < len(b) && j < len(a); {
			if b[i] < a[j] {
				missingFromA = append(missingFromA, b[i])
				i += 1
			} else if b[i] > a[j] {
				j += 1
			} else {
				i += 1
				j += 1
			}
		}
		if j < len(a) {
			missingFromB = append(missingFromB, a[j:]...)
		}
	}
	return
}

// Complements determine the complements of both strings arrays.
func Complements(as, bs []string) (missingFromA, missingFromB []string) {
	setA := make(map[string]bool)
	for _, a := range as {
		setA[a] = true
	}

	setB := make(map[string]bool)
	for _, b := range bs {
		setB[b] = true
	}

	for _, a := range as {
		if _, ok := setB[a]; !ok {
			missingFromB = append(missingFromB, a)
		}
	}

	for _, b := range bs {
		if _, ok := setA[b]; !ok {
			missingFromA = append(missingFromA, b)
		}
	}
	return
}
