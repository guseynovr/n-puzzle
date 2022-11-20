package puzzle

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Parse parses Puzzle struct from filename.
// If filename is an empty string, reads from stdin.
func Parse(filename string) (*Puzzle, error) {
	var file *os.File
	var err error
	if filename == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	var size int
	var tiles []int
	for scanner.Scan() {
		line := rmCommentAndTrim(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if size == 0 {
			size, err = strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("parsing size: %w", err)
			}
			if size < 3 {
				return nil, fmt.Errorf("size must be > 2, got %d", size)
			}
			tiles = make([]int, 0, size*size)
			continue
		}
		newtiles, err := scanTiles(line)
		if err != nil {
			return nil, fmt.Errorf("scan tiles: %w", err)
		}
		tiles = append(tiles, newtiles...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("scanning file: ", err)
	}
	if size == 0 {
		return nil, fmt.Errorf("empty input")
	}
	if len(tiles) != size*size {
		return nil, fmt.Errorf("incorrect number of tiles, expected %d, got %d",
			size*size, len(tiles))
	}

	x, y, err := validateTiles(tiles, size)
	if err != nil {
		return nil, fmt.Errorf("validatetiles: %w", err)
	}
	return newPuzzle(size, x, y, to2D(tiles, size)), nil
}

func validateTiles(tiles []int, size int) (int, int, error) {
	unique := make(map[int]struct{})
	keys := make([]int, len(tiles))
	x, y := 0, 0
	for i, v := range tiles {
		if _, exists := unique[v]; exists {
			return 0, 0, fmt.Errorf("tile values must be unique")
		}
		unique[v] = struct{}{}
		keys[i] = v
		if v == 0 {
			x = i % size
			y = i / size
		}
	}
	sort.Ints(keys)
	if keys[0] != 0 {
		return 0, 0, fmt.Errorf("must have one empty tile (with 0 value)")
	}
	for i := 0; i < len(keys); i++ {
		if i != keys[i] {
			return 0, 0, fmt.Errorf("invalid tile values")
		}
	}
	return x, y, nil
}

func to2D(slice []int, size int) [][]int {
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}
	for i, val := range slice {
		result[i/size][i%size] = val
	}
	return result
}

// scanTiles extracts a slice of ints from space separated string of numbers.
func scanTiles(line string) ([]int, error) {
	var err error
	fields := strings.Fields(line)
	tiles := make([]int, len(fields))
	for i, f := range fields {
		tiles[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return tiles, nil
}

func rmCommentAndTrim(line string) string {
	if strings.ContainsRune(line, '#') {
		line = line[:strings.IndexRune(line, '#')]
	}
	return strings.TrimSpace(line)
}
