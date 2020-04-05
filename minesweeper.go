package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type square struct {
	bomb    bool
	flag    bool
	visible bool
	count   *int
}
type row []*square
type board []row

var gameWon *bool

func makeBoard() board {
	b := board{}

	// to do: add ability to choose board size
	for i := 0; i < 4; i++ {
		r := row{}
		for j := 0; j < 7; j++ {
			c := rand.Intn(100)
			s := square{
				visible: false,
			}
			// to do: re-seed randomness
			if c > 85 {
				s.bomb = true
			}
			// use make and index instead of append
			r = append(r, &s)
		}
		b = append(b, r)
	}
	for i, row := range b {
		for j, square := range row {
			square.setCount(b, i, j)
		}
	}
	return b
}

// Calculate bomb proximity count for a square
func (s *square) setCount(b board, i, j int) {
	count := 0
	s.count = &count

	leftedge := i == 0
	rightedge := i == len(b)-1
	topedge := j == 0
	bottomedge := j == len(b[i])-1

	// pretty ugly, any way to make this nicer?
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

func (b board) clickSquare(i, j int) error {
	if i >= len(b) || j >= len(b[0]) {
		return errors.New("Square coordinates out of range")
	}
	// lose game
	if b[i][j].bomb {
		w := false
		gameWon = &w
		return nil
	}
	b[i][j].visible = true
	return nil
}

func (b board) flagSquare(i, j int) error {
	if i >= len(b) || j >= len(b[0]) {
		return errors.New("Square coordinates out of range")
	}
	b[i][j].flag = true
	return nil
}

func (b board) revealBoard() {
	for _, row := range b {
		for _, square := range row {
			square.visible = true
		}
	}
}

func (b board) printBoard() error {

	// mark to false if board has not been fully/properly revealed
	gameComplete := true

	fmt.Println(strings.Repeat("-", 4*len(b[0])+1))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			s := b[i][j]
			if s.count == nil {
				return fmt.Errorf("missing count for square: %d, %d", i, j)
			}
			var val string
			switch {
			case s.flag:
				val = "F"
				if !s.bomb {
					// cannot win, bomb incorrectly flagged
					gameComplete = false
				}
			case !s.visible:
				val = "."
				// game incomplete, non-bombs still uncovered
				// (flags on bombs are optional)
				if !s.bomb {
					gameComplete = false
				}
			case s.bomb:
				val = "B"
			default:
				val = strconv.Itoa(*s.count)
			}
			fmt.Printf("| %v ", val)
		}
		fmt.Printf("|\n%s \n", strings.Repeat("-", 4*len(b[i])+1))
		// all non-bombs have been revealed
		if gameComplete {
			w := true
			gameWon = &w
		}
	}
	return nil
}

func main() {
	b := makeBoard()
	if err := b.printBoard(); err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// add nums along edges for easier gameplay(?)
	// stats as you go (how many bombs flagged, etc.)

	reader := bufio.NewReader(os.Stdin)

	for gameWon == nil {

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)

		chars := strings.Split(text, ",")
		// to do: handle incorrect input
		x, err := strconv.Atoi(chars[0])
		y, err := strconv.Atoi(chars[1])
		if err != nil {
			fmt.Println("Invalid move, must provide number")
		}
		if len(chars) > 2 && strings.ToLower(chars[2]) == "f" {
			if err := b.flagSquare(y, x); err != nil {
				fmt.Printf("Error: %v\n\n", err)
			}
		} else {
			if err := b.clickSquare(y, x); err != nil {
				fmt.Printf("Error: %v\n\n", err)
			}
		}

		if err := b.printBoard(); err != nil {
			fmt.Printf("Error: %v\n\n", err)
			break
		}
		// determined during print
		if gameWon != nil {
			break
		}

	}

	if *gameWon {
		fmt.Println("\nWOOO YOU WIN!")
	} else {
		fmt.Println("\nBOOM! GAME OVER")
	}
	b.revealBoard()
	b.printBoard()

}
