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
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	filename := os.Args[1]
	p, err := puzzle.Parse(filename)
	if err != nil {
		log.Fatal(err)
	}
	// p := puzzle.Random(2)
	fmt.Println(p)
	if !p.IsSolvable() {
		fmt.Println("Puzzle is unsolvable")
		os.Exit(0)
	}
}

// TODO: func processOptions() (size int, )
