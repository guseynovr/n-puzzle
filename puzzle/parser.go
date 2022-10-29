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
func Parse(filename string) (*Puzzle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var size int
	var cells []int
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
			cells = make([]int, 0, size*size)
			continue
		}
		newCells, err := scanCells(line)
		if err != nil {
			return nil, fmt.Errorf("scan cells: %w", err)
		}
		cells = append(cells, newCells...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("scanning file: ", err)
	}
	if len(cells) != size*size {
		return nil, fmt.Errorf("incorrect number of cells, expected %d, got %d",
			size*size, len(cells))
	}

	result := Puzzle{
		size:  size,
		cells: to2D(cells, size)}
	if err := result.validateCells(cells); err != nil {
		return nil, fmt.Errorf("validateCells: %w", err)
	}
	return &result, nil
}

func (p *Puzzle) validateCells(cells []int) error {
	unique := make(map[int]struct{})
	keys := make([]int, len(cells))
	for i, v := range cells {
		if _, exists := unique[v]; exists {
			return fmt.Errorf("cell values must be unique")
		}
		unique[v] = struct{}{}
		keys[i] = v
		if v == 0 {
			p.emptyX = i % p.size
			p.emptyY = i / p.size
		}
	}
	sort.Ints(keys)
	if keys[0] != 0 {
		return fmt.Errorf("must have one empty cell (with 0 value)")
	}
	for i := 0; i < len(keys); i++ {
		if i != keys[i] {
			return fmt.Errorf("invalid cell values")
		}
	}
	return nil
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

// scanCells extracts a slice of ints from space separated string of numbers.
func scanCells(line string) ([]int, error) {
	var err error
	fields := strings.Fields(line)
	cells := make([]int, len(fields))
	for i, f := range fields {
		cells[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return cells, nil
}

func rmCommentAndTrim(line string) string {
	if strings.ContainsRune(line, '#') {
		line = line[:strings.IndexRune(line, '#')]
	}
	return strings.TrimSpace(line)
}
