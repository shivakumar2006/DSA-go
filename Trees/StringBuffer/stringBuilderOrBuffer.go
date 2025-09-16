// collecting lines from files -> using strings.Builder
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "/users/shivakumar/projects/DSA-go/Trees/AVL.go"

	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file : ", err)
		return
	}

	// print the collected output
	fmt.Println("File content : ")
	fmt.Println(sb.String())
}

// // now mixing byte and string -> usinf bytes.Buffer

// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func main() {
// 	var b bytes.Buffer

// 	b.Write([]byte("Hello "))
// 	b.WriteByte('G')
// 	b.WriteByte('o')
// 	fmt.Println(b.String())
// }

// // strings builder, Strings.Builder

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// JSON like output
// 	var sb strings.Builder
// 	sb.WriteString("{")
// 	sb.WriteString(`"name":"Shiva",`)
// 	sb.WriteString(`"age":"19"`)
// 	sb.WriteString("}")
// 	fmt.Println(sb.String())

// 	// // Pre-Allocating buffer size
// 	// var sb strings.Builder
// 	// sb.Grow(100) // reserve 100 Bytes

// 	// sb.WriteString("Hello ")
// 	// sb.WriteString("World!")
// 	// fmt.Println(sb.String())

// 	// var sb strings.Builder

// 	// for i := 0; i <= 5; i++ {
// 	// 	sb.WriteString(fmt.Sprintf("Item %d\n", i))
// 	// }

// 	// result := sb.String()
// 	// fmt.Println(result)

// 	// var builder strings.Builder
// 	// builder.WriteString("Hello")
// 	// builder.WriteString(" Kumar")
// 	// fmt.Println(builder.String())
// }

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
