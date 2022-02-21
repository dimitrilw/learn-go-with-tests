package arraynslice

func Sum(n []int) (sum int) {
	for _, v := range n {
		sum += v
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, s := range numbersToSum {
		sums = append(sums, Sum(s))
	}
	return
}
