package main

func Sum(n []int) (val int) {
	for _, item := range n {
		val += item
	}

	return
}

func SumAll(n ...[]int) (sums []int) {
	for _, numbers := range n {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(n ...[]int) (sums []int) {
	for _, numbers := range n {
		if len(numbers) == 0 {
			sums = append(sums, Sum(numbers))

		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
