// https://go-tour-jp.appspot.com/flowcontrol/12

package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}
