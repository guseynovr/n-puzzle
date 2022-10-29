package main

import (
	"fmt"
	"log"
	"os"

	"n-puzzle/puzzle"
)

func main() {
	//TODO: choose heuristic
	//TODO: get file from args
	//TODO: generate random if filepath not given
	p, err := puzzle.Parse("examples/4-3.nok.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	if !p.IsSolvable() {
		fmt.Println("Puzzle is unsolvable")
		os.Exit(0)
	}
}
