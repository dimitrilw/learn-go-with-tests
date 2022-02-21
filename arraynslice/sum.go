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

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, s := range numbersToSum {
		if len(s) < 2 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(s[1:]))
		}
	}
	return
}
