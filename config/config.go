package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"npuzzle/algorithm"
	"npuzzle/puzzle"
)

type Config struct {
	random    size
	file      string
	Pause     time.Duration
	Blocks    bool
	Heuristic heuristic
	Debug     bool
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
			Desc: "manhattan",
		}, "`heuristics` to be used: Manhattan, Euclidean, Diagonal, Tiles")
	fs.DurationVar(&cfg.Pause, "p", time.Millisecond*200,
		"pause between steps in animation in `milliseconds`")
	fs.BoolVar(&cfg.Blocks, "b", false, "show blocks")
	fs.BoolVar(&cfg.Debug, "debug", false, "use debug pauses and logs")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, err
	}
	if cfg.random != 0 && cfg.file != "" {
		return nil, fmt.Errorf("-f and -r are mutually exclusive")
	}
	if cfg.Debug && cfg.file == "" {
		return nil, fmt.Errorf("debug can't be used when reading from stdin")
	}
	return &cfg, nil
}

func (c Config) Forge() (*puzzle.Puzzle, error) {
	var res *puzzle.Puzzle
	var err error
	if c.random > 0 {
		res = puzzle.Random(int(c.random))
	} else {
		res, err = puzzle.Parse(c.file)
		if err != nil {
			return nil, err
		}
	}
	res.Blocks = c.Blocks
	return res, nil
}
