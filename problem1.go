// sum of multiple of 3 and 5 below 1000
package EulerProject

//3 U 5= either of 3 and either of 5

func problem1(number int) int {
	var sum int = 0
	for i := 1; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
