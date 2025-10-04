package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 1, 4, 3, 2, 6}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				arr[i], arr[j] = arr[j], arr[i]
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		swapped := false
// 		for j := 1; j < len(arr); j++ {
// 			if arr[j] < arr[j-1] {
// 				temp := arr[j]
// 				arr[j] = arr[j-1]
// 				arr[j-1] = temp
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return arr
// }
