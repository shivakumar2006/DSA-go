// strings builder, Strings.Builder

package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" Kumar")
	fmt.Println(builder.String())
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
