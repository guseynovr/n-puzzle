package config

import (
	"fmt"
	"strconv"
	"strings"
)

type size int

func (s *size) UnmarshalText(text []byte) error {
	n, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}
	if n < 3 {
		return fmt.Errorf("size cannot be less than 3")
	}
	*s = size(n)
	return nil
}

func (s size) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprint(int(s))), nil
}

type heuristic string

func (h *heuristic) UnmarshalText(text []byte) error {
	*h = heuristic(strings.ToLower(string(text)))
	switch *h {
	case "manhattan", "diagonal", "euclidean":
	default:
		return fmt.Errorf("unsupported heuristic: %s", *h)
	}
	return nil
}

func (h heuristic) MarshalText() ([]byte, error) {
	return []byte(h), nil
}
