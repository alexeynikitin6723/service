package wonderful

func Solution(A []int) int {
	seq := map[int]int{}
	for i := 0; i < len(A); i++ {
		seq[A[i]]++
	}
	for x, y := range seq {
		if y%2 == 1 {
			return x
		}
	}
	return 0
}
