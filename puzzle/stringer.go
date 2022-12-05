package puzzle

import (
	"fmt"
	"strings"
)

const (
	red    = "\033[%s1m"
	green  = "\033[%s2m"
	yellow = "\033[%s3m"
	blue   = "\033[%s4m"
	reset  = "\033[0m"

	regular = "0;3"
	blocked = "4"
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
	width := len(fmt.Sprint(p.Size*p.Size - 1 /* , p.Tiles[0][0].Target */))
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
func (ps *puzzleStringer) writeRow(row []Tile) {
	ps.sb.WriteRune('│')
	for j, tile := range row {
		ps.writeValue(tile)
		if j < ps.size-1 {
			ps.sb.WriteRune('│')
		}
	}
	ps.sb.WriteString("│\n")
}

func (ps *puzzleStringer) writeValue(tile Tile) {
	color := red
	tileType := regular

	if tile.Relevant {
		color = green
	}
	if tile.Locked {
		tileType = blocked
	}
	if tile.Value == 0 {
		color = yellow
	}
	color = fmt.Sprintf(color, tileType)
	// value := fmt.Sprintf("%d%v", tile.Value, tile.Target)
	value := fmt.Sprintf("%d", tile.Value)
	ps.sb.WriteString(fmt.Sprintf("%s%*s%s", color, ps.width, value, reset))
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
