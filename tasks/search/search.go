package search

func Solution(A []int) int {
	n := len(A) + 1
	sum := (1 + n) * n / 2
	sumArr := 0
	for i := 0; i < len(A); i++ {
		sumArr += A[i]
	}
	return sum - sumArr
}

// func contains(elem int, A []int) bool {
// 	for _, i := range A {
// 		if elem == i {
// 			return true
// 		}
// 	}
// 	return false
// }
