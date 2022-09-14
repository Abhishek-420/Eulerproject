package EulerProject

// sum of even fibonacci numbers less than 400000

func Problem2() int {
	const max = 400000
	first := 0
	second := 1
	sum := 0
	for second < max {
		if second%2 == 0 {
			sum += second
		}
		temp := first + second
		first = second
		second = temp
	}
	return sum
}
