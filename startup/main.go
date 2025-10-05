package main

import (
	"fmt"

	"github.com/gobuffalo/flect"
)

func main() {
	s := "hello, world!"
	d := flect.Dasherize(s)
	fmt.Println(d)
	fmt.Println(flect.Singularize("people"))
}
