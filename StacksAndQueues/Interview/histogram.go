// Largest Rectangle in histogram...
// Google, Amazon Facebook question.

package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("Largest Rectangle area : ", largestRectangleArea(heights)) // 10
}

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[topIndex]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

// func largestRectangeArea(heights []int) int {
// 	stack := []int{}
// 	maxArea := 0
// 	n := len(heights)

// 	for i := 0; i <= n; i++ {
// 		h := 0
// 		if i < n {
// 			h = heights[i]
// 		}

// 		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
// 			// pop the height
// 			topIndex := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			height := heights[topIndex]
// 			width := i
// 			if len(stack) > 0 {
// 				width = i - stack[len(stack)-1] - 1
// 			}
// 			area := height * width
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 		stack = append(stack, i)
// 	}
// 	return maxArea
// }

// func largestRectangeArea(heights []int) int {
// 	stack := []int{} // store indices
// 	maxArea := 0
// 	n := len(heights)

// 	for i := 0; i <= n; i++ {
// 		// treat i == n as height == 0
// 		h := 0
// 		if i < n {
// 			h = heights[i]
// 		}

// 		// maintain increasing stack
// 		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
// 			// pop the height
// 			topIndex := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			height := heights[topIndex]
// 			width := i
// 			if len(stack) > 0 {
// 				width = i - stack[len(stack)-1] - 1
// 			}
// 			area := height * width
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 		stack = append(stack, i)
// 	}
// 	return maxArea
// }
