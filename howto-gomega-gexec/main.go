package main

import (
	"flag"
	"fmt"
)

func main() {

	a := flag.String("first", "hoge", "a is string type")
	flag.Parse()
	fmt.Println(*a)
}
