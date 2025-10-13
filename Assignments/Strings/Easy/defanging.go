// Defanging an IP Address
// https://leetcode.com/problems/defanging-an-ip-address/description/

package main

import (
	"fmt"
	"strings"
)

func defangingIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	fmt.Println(defangingIPaddr("1.1.1.1"))
	fmt.Println(defangingIPaddr("255.100.50.0"))
}
