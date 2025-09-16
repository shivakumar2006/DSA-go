// strings builder, Strings.Builder

package main

import (
	"fmt"
	"strings"
)

func main() {
	// JSON like output
	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString(`"name":"Shiva",`)
	sb.WriteString(`"age":"19"`)
	sb.WriteString("}")
	fmt.Println(sb.String())

	// // Pre-Allocating buffer size
	// var sb strings.Builder
	// sb.Grow(100) // reserve 100 Bytes

	// sb.WriteString("Hello ")
	// sb.WriteString("World!")
	// fmt.Println(sb.String())

	// var sb strings.Builder

	// for i := 0; i <= 5; i++ {
	// 	sb.WriteString(fmt.Sprintf("Item %d\n", i))
	// }

	// result := sb.String()
	// fmt.Println(result)

	// var builder strings.Builder
	// builder.WriteString("Hello")
	// builder.WriteString(" Kumar")
	// fmt.Println(builder.String())
}

// // string buffer bytes.Buffer

// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func main() {
// 	var buf bytes.Buffer
// 	buf.WriteString("Hello")
// 	buf.WriteString(" World")
// 	fmt.Println(buf.String()) // "Hello World"
// }
