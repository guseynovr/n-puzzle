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
	fmt.Printf("Start:\n%s\n", p)
	if !p.IsSolvable() {
		fmt.Println("Puzzle is unsolvable")
		return
	}
	if p.IsSolved() {
		fmt.Println("Solved!")
		return
	}

	s := algorithm.Solver{
		P:     p,
		H:     cfg.Heuristic.F,
		Debug: false,
	}
	stats := s.Solve()

	fmt.Println("Finished!")
	fmt.Println(stats.Path[len(stats.Path)-1])
	fmt.Println("Heuristics used:", cfg.Heuristic.Desc)
	fmt.Println(stats)
}
