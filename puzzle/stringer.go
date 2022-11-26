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

/*
String prints puzzle as a table.

	┌─┬─┬─┐
	│0│1│2│
	├─┼─┼─┤
	│3│4│5│
	├─┼─┼─┤
	│6│7│8│
	└─┴─┴─┘
*/
func (p Puzzle) String() string {
	width := len(fmt.Sprint(p.Size*p.Size - 1))
	horizontal := strings.Repeat("─", width)

	ps := puzzleStringer{
		sb:         strings.Builder{},
		width:      width,
		horizontal: horizontal,
		size:       p.Size,
	}
	ps.writeTopLine()
	for i, row := range p.Tiles {
		ps.writeRow(row)
		if i < p.Size-1 {
			ps.writeMiddleLine()
		}
	}
	ps.writeBottomLine()
	return ps.sb.String()
}

// "│x│x│\n"
func (ps *puzzleStringer) writeRow(row []int) {
	ps.sb.WriteRune('│')
	for j, tile := range row {
		ps.sb.WriteString(fmt.Sprintf("%*d", ps.width, tile))
		if j < ps.size-1 {
			ps.sb.WriteRune('│')
		}
	}
	ps.sb.WriteString("│\n")
}

// "┌─┬─┐\n"
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

// "├─┼─┤\n"
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

// "└─┴─┘"
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
