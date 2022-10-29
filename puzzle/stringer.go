package puzzle

import (
	"fmt"
	"strings"
)

type puzzleStringer struct {
	sb         strings.Builder
	width      int
	horizontal string
	size       int
}

/*	┌─┬─┬─┐
	│ │ │ │
	├─┼─┼─┤
	│ │ │ │
	├─┼─┼─┤
	│ │ │ │
	└─┴─┴─┘ */
func (p *Puzzle) String() string {
	width := len(fmt.Sprint(p.size*p.size - 1))
	horizontal := strings.Repeat("─", width)

	ps := puzzleStringer{
		sb:         strings.Builder{},
		width:      width,
		horizontal: horizontal,
		size:       p.size,
	}
	ps.writeTopLine()
	for i, row := range p.cells {
		ps.writeRow(row)
		if i < p.size-1 {
			ps.writeMiddleLine()
		}
	}
	ps.writeBottomLine()
	return ps.sb.String()
}

func (ps *puzzleStringer) writeRow(row []int) {
	ps.sb.WriteRune('│')
	for j, cell := range row {
		ps.sb.WriteString(fmt.Sprintf("%*d", ps.width, cell))
		if j < ps.size-1 {
			ps.sb.WriteRune('│')
		}
	}
	ps.sb.WriteString("│\n")
}

func (ps *puzzleStringer) writeTopLine() {
	ps.sb.WriteRune('┌')
	for i := 0; i < ps.size; i++ {
		ps.sb.WriteString(ps.horizontal)
		if i < ps.size-1 {
			ps.sb.WriteRune('┬')
		}
	}
	ps.sb.WriteString("┐\n")
}

func (ps *puzzleStringer) writeMiddleLine() {
	ps.sb.WriteRune('├')
	for i := 0; i < ps.size; i++ {
		ps.sb.WriteString(ps.horizontal)
		if i < ps.size-1 {
			ps.sb.WriteRune('┼')
		}
	}
	ps.sb.WriteString("┤\n")
}

func (ps *puzzleStringer) writeBottomLine() {
	ps.sb.WriteRune('└')
	for i := 0; i < ps.size; i++ {
		ps.sb.WriteString(ps.horizontal)
		if i < ps.size-1 {
			ps.sb.WriteRune('┴')
		}
	}
	ps.sb.WriteRune('┘')
}
