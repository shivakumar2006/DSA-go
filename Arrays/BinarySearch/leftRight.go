package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 0, 1, 0, 1, 0, 0}
	res := split(arr)
	fmt.Println(res)
}

func split(arr []int) []int {
	start, end := 0, len(arr)-1
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			res[start] = 0
			start++
		} else {
			res[end] = 1
			end--
		}
	}
	return res
}

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"encoding/json"
// 	"math/rand"
// 	"net/http"
// 	"time"
// )

// type ApiResponse struct {
// 	Message string `json:"message"`
// 	Value   int    `json:"value"`
// 	Time    string `json:"time"`
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	rand.Seed(time.Now().UnixNano())

// 	response := ApiResponse{
// 		Message: "Here is your random value",
// 		Value:   rand.Intn(100), // random number between 0-99
// 		Time:    time.Now().Format(time.RFC3339),
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func main() {
// 	http.HandleFunc("/", handler)

// 	port := ":8080"
// 	println("Server is running on http://localhost" + port)
// 	http.ListenAndServe(port, nil)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 0, 1, 0, 1, 0}
// 	res := split(arr)
// 	fmt.Println(res)
// }

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 0, 1, 0, 1, 0}
// 	res := split(arr)
// 	fmt.Println(res)
// }

// func split(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			res[start] = 0
// 			start++
// 		} else {
// 			res[end] = 1
// 			end--
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 0, 1, 0, 1, 1}
// 	res := leftRight(arr)
// 	fmt.Println(res)
// }

// func leftRight(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	result := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			result[start] = 0
// 			start++
// 		} else {
// 			result[end] = 1
// 			end--
// 		}
// 	}
// 	return result
// }
