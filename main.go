package main

import (
	"fmt"
	"log"

	"npuzzle/algorithm"
	"npuzzle/config"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	p, err := cfg.Forge()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	if !p.IsSolvable() {
		fmt.Println("Puzzle is unsolvable")
		return
	}
	if p.IsSolved() {
		fmt.Println("Solved!")
		return
	}
	stats := algorithm.AStar(*p, cfg.Heuristic.F)
	_ = stats
}
