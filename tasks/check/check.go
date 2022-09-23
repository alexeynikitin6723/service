package check

func Solution(A []int) int {
	seq := map[int]int{}
	max := 0
	for i := 0; i < len(A); i++ {
		if seq[A[i]] != 0 {
			return 0
		} else {
			seq[A[i]]++
			if max < A[i] {
				max = A[i]
			}
		}
	}
	if (max) == len(A) {
		return 1
	} else {
		return 0
	}
}

// func Solution(A []int) int {
// 	sort.Ints(A)
// 	max := A[len(A)-1]
// 	if max == len(A) {
// 		return 1
// 	}
// 	return 0
// }
