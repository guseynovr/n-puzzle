package config

import (
	"flag"
	"fmt"
	"os"

	"npuzzle/algorithm"
	"npuzzle/puzzle"
)

type Config struct {
	random    size
	file      string
	Heuristic heuristic
}

func Parse() (*Config, error) {
	var cfg Config
	fs := flag.NewFlagSet("npuzzle", flag.ExitOnError)
	fs.TextVar(&cfg.random, "r", size(0),
		"generate random puzzle of given `size`")
	fs.StringVar(&cfg.file, "f", "",
		"`path` to the file with a starting board")
	fs.TextVar(&cfg.Heuristic, "he",
		heuristic{
			F:    algorithm.Manhattan,
			desc: "manthattan",
		}, "`heuristics` to be used: Manhattan, Euclidean, Out-of-place")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, err
	}
	if cfg.random != 0 && cfg.file != "" {
		return nil, fmt.Errorf("-f and -r are mutually exclusive")
	}
	return &cfg, nil
}

func (c Config) Forge() (*puzzle.Puzzle, error) {
	if c.random > 0 {
		return puzzle.Random(int(c.random)), nil
	}
	return puzzle.Parse(c.file)
}
