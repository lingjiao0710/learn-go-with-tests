package main

import "fmt"

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersTosum ...[]int) (sums []int) {
	//lengthOfNumbers := len(numbersTosum)
	//sums = make([]int, lengthOfNumbers)
	//var sums []int

	for _, numbers := range numbersTosum {
		//sums[i] = Sum(numbers)
		if 0 == len(numbers) {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers))
		}

	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {

		if 0 == len(numbers) {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(arr))
}
