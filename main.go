package main

import (
	"fmt"
	"log"
	"time"

	"npuzzle/algorithm"
	"npuzzle/config"
	"npuzzle/puzzle"
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
	p.MakeAllIrrelevant()
	fmt.Printf("Start:\n%s\n", p)
	p.MakeAllRelevant()
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
		Debug: cfg.Debug,
		ByH:   true,
	}
	stats := s.Solve()

	fmt.Println("Finished!")
	fmt.Println(stats.Path[len(stats.Path)-1])
	fmt.Println("Heuristics used:", cfg.Heuristic.Desc)
	fmt.Println(stats)

	ch := make(chan struct{})
	go animatePath(ch, p.Size, stats.Path, cfg.Pause)
	<-ch
}

func animatePath(ch chan struct{}, size int, path []puzzle.Puzzle,
	pause time.Duration) {
	nlines := size*2 + 2

	fmt.Printf("\033[%dS", nlines) // scroll up to make room for output
	fmt.Printf("\033[%dA", nlines) // move cursor back up
	fmt.Print("\033[s")            // save current cursor position

	for i, step := range path {
		fmt.Print("\033[u") // restore saved cursor position
		fmt.Println(step, "Step:", i)
		time.Sleep(pause)
	}
	ch <- struct{}{}
}
