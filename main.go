package main

import (
	"fmt"

	"github.com/code-chut-xiu/go-design-patterns/creational/singleton"
)

func main() {
	testSingleton()
}

func testSingleton() {
	s1 := singleton.New()

	s1["this"] = "that"

	s2 := singleton.New()

	fmt.Println("This is ", s2["this"])
}
