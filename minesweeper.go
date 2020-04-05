package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type square struct {
	bomb    bool
	visible bool
	count   *int
}
type row []*square
type board []row

// Calculate bomb proximity count for a square
func (s *square) setCount(b board, i, j int) {
	count := 0
	s.count = &count

	leftedge := i == 0
	rightedge := i == len(b)-1
	topedge := j == 0
	bottomedge := j == len(b[i])-1

	if !leftedge {
		if b[i-1][j].bomb {
			*s.count++
		}
		if !topedge {
			if b[i-1][j-1].bomb {
				*s.count++
			}
		}
		if !bottomedge {
			if b[i-1][j+1].bomb {
				*s.count++
			}
		}
	}
	if !rightedge {
		if b[i+1][j].bomb {
			*s.count++
		}
		if !topedge {
			if b[i+1][j-1].bomb {
				*s.count++
			}
		}
		if !bottomedge {
			if b[i+1][j+1].bomb {
				*s.count++
			}
		}
	}
	if !topedge {
		if b[i][j-1].bomb {
			*s.count++
		}
	}
	if !bottomedge {
		if b[i][j+1].bomb {
			*s.count++
		}
	}
}

func printBoard(b board) error {
	fmt.Println(strings.Repeat("-", 4*len(b[0])+1))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			s := b[i][j]
			if s.count == nil {
				return fmt.Errorf("missing count for square: %d, %d", i, j)
			}
			var val string
			switch {
			case !s.visible:
				val = "X"
			case s.bomb:
				val = "B"
			default:
				val = strconv.Itoa(*s.count)
			}
			fmt.Printf("| %v ", val)
		}
		fmt.Printf("|\n%s \n", strings.Repeat("-", 4*len(b[i])+1))
	}
	return nil
}

func makeBoard() *board {
	b := board{}

	// to do: add ability to choose board size
	for i := 0; i < 5; i++ {
		r := row{}
		for j := 0; j < 7; j++ {
			c := rand.Intn(100)
			s := square{
				visible: false,
			}
			// to do: re-seed randomness
			if c > 70 {
				s.bomb = true
			}
			r = append(r, &s)
		}
		b = append(b, r)
	}
	for i, row := range b {
		for j, square := range row {
			square.setCount(b, i, j)
		}
	}
	return &b
}

func main() {
	b := makeBoard()
	if err := printBoard(*b); err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}
	// add gameplay next
}
