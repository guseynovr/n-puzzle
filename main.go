package main

import (
	"fmt"
	"log"

	"./parser"
)

func main() {
	//TODO: choose heuristic
	//TODO: get file from args
	//TODO: generate random if filepath not given
	p, err := parser.Parse("examples/3-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
