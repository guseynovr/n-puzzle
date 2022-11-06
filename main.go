package main

import (
	"fmt"
	"log"
	"os"

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
		os.Exit(0)
	}
}
