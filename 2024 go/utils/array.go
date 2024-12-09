package utils

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func SumList(s []int) uint64 {
	var sum uint64 = 0

	for _, e := range s {
		sum += uint64(e)
	}

	return sum
}

func Reverse(s []int) []int {
	size := len(s)
	opposite := make([]int, size)

	for i, e := range s {
		opposite[size - 1 - i] = e
	}

	return opposite
}