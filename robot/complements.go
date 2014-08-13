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

// Complements determine the complements of both arrays.
func Complements(a, b []int64) (missingFromA, missingFromB []int64) {
	sort.Sort(Int64Slice(a))
	sort.Sort(Int64Slice(b))
	{
		var j int
		for _, i := range a {
			if (j + 1) > len(b) {
				break
			}
			if i < b[j] {
				missingFromB = append(missingFromB, i)
			} else {
				j += 1
			}
		}
		if j < len(b) {
			missingFromA = append(missingFromA, b[j:]...)
		}
	}
	{
		var j int
		for _, i := range b {
			if (j + 1) > len(a) {
				break
			}
			if i < a[j] {
				missingFromA = append(missingFromA, i)
			} else {
				j += 1
			}
		}
		if j < len(a) {
			missingFromB = append(missingFromB, a[j:]...)
		}
	}
	return
}
