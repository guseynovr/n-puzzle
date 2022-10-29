package main

import (
	"fmt"
	"log"

	"n-puzzle/puzzle"
)

func main() {
	//TODO: choose heuristic
	//TODO: get file from args
	//TODO: generate random if filepath not given
	p, err := puzzle.Parse("examples/4-1_tabs.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
