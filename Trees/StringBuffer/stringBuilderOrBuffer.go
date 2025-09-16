// string buffer bytes.Buffer

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteString(" World")
	fmt.Println(buf.String()) // "Hello World"
}
