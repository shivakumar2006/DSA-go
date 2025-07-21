package main

import "fmt"

func main() {
	arr := []int{-2, 5, 69, -100, 45}
	tar := 69
	res := search(arr, tar)
	if res != -1 {
		fmt.Printf("THe target element is found : %d\n", res)
	} else {
		fmt.Println("Target not found")
	}
}

func search(arr []int, tar int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == tar {
			return i
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, 69, -100, 45}
// 	tar := -100
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element is found : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	for i, val := range arr {
// 		if val == tar {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, 69, -100, 45}
// 	tar := 69
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element is found : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	for i, val := range arr {
// 		if val == tar {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, -3, 69, -100, 45}
// 	tar := 69
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element is found : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == tar {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, -3, 7, 69, -100, 45}
// 	target := 69

// 	res := linearSearch(arr, target)
// 	if res != -1 {
// 		fmt.Printf("Element is found at the index of : %d\n", res)
// 	} else {
// 		fmt.Println("Element is not found...")
// 	}
// }

// func linearSearch(arr []int, target int) int {
// 	for i := 0; i <= len(arr); i++ {
// 		if arr[i] == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, -3, 7, 69, -100, 45}

// 	target := 69
// 	res := linearSearch(arr, target)
// 	if res != -1 {
// 		fmt.Printf("Element is found at the index of %d\n", res)
// 	} else {
// 		fmt.Println("Element is not found")
// 	}
// }

// func linearSearch(arr []int, target int) int {
// 	for i := 0; i <= len(arr); i++ {
// 		if arr[i] == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, -3, 7, 69, -100, 45}
// 	target := -100

// 	result := linearSearch(arr, target)
// 	if result != -1 {
// 		fmt.Printf("Element is found at the index of %d\n", result)
// 	} else {
// 		fmt.Println("Element is not found")
// 	}
// }

// func linearSearch(arr []int, target int) int {
// 	for i := 0; i <= len(arr); i++ {
// 		if arr[i] == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func linearSearch(arr []int , target int) int {
// 	for i, value := range arr {
// 		if value == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// BOOL linear search
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 3, 7, 69, 100, 45}
// 	target := 69

// 	result := linearSearch(arr, target)
// 	if result != false {
// 		fmt.Println("Element is found at the index \n", result)
// 	} else {
// 		fmt.Println("Element is not found")
// 	}
// }

// func linearSearch(arr []int, target int) bool {
// 	for i := 0; i <= len(arr); i++ {
// 		if arr[i] == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 3, 7, 69, 100, 45}
// 	target := 69

// 	result := linearSearch(arr, target)
// 	if result != -1 {
// 		fmt.Printf("Element found at index %d\n", result)
// 	} else {
// 		fmt.Println("Element not found")
// 	}
// }

// func linearSearch(arr []int, target int) int {
// 	for i, value := range arr {
// 		if value == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	array := []int{1, 4, 5, 6, 3, 7, 10, 69}
// 	target := 7

// 	result := linearSearch(array, target)

// 	if result != -1 {
// 		fmt.Printf("Element found at index %d\n", result)
// 	} else {
// 		fmt.Println("Element not found")
// 	}
// }

// func linearSearch(array []int, target int) int {
// 	for i, value := range array {
// 		if value == target {
// 			return i
// 		}
// 	}
// 	return -1
// }
