// Average Salary Excluding the Minimum and Maximum Salary
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/description/

package main

import "fmt"

func average(salary []int) float64 {
	max, min := salary[0], salary[0]
	for i := 0; i < len(salary); i++ {
		if salary[i] > max {
			max = salary[i]
		} else if salary[i] < min {
			min = salary[i]
		}
	}

	sum := 0.0
	count := 0
	for i := 0; i < len(salary); i++ {
		if salary[i] != max && salary[i] != min {
			sum += float64(salary[i])
			count++
		}
	}

	return sum / float64(count)
}

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))
}
