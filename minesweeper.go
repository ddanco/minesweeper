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
type row []square
type board []row

func (s square) setCount(b board, i, j int) square {
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
	return s
}

func printBoard(b board) {
	fmt.Println(strings.Repeat("-", 4*len(b)+1))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			s := b[i][j]
			if s.count == nil {
				s = s.setCount(b, i, j)
			}
			var val string
			if s.bomb {
				val = "B"
			} else {
				val = strconv.Itoa(*s.count)
			}
			fmt.Printf("| %v ", val)
		}
		fmt.Printf("|\n%s \n", strings.Repeat("-", 4*len(b)+1))
	}
}

func makeBoard() board {
	b := board{}

	// add ability to choose board size
	for i := 0; i < 7; i++ {
		r := row{}
		for j := 0; j < 7; j++ {
			c := rand.Intn(100)
			s := square{
				visible: false,
			}
			// add ability to choose number of bombs
			// also re-seed randomness
			if c > 70 {
				s.bomb = true
			}
			r = append(r, s)
		}
		b = append(b, r)
	}
	return b
}

func main() {
	b := makeBoard()
	printBoard(b)
	// add gameplay next
}
