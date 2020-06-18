package snakesladders

import (
	"math/rand"
	"time"
)

// Main components of the game
//  -- player and computer
//     position in which the player is at
//  -- the board - a single array from 0 - 99 with values of cell 1 - 100
//  -- snakes
//     list of snakes with ( head and tail position)
//  -- ladders
//     list of ladders with ( up and down position)
//

type snake struct {
	head int
	tail int
}

type ladder struct {
	top    int
	bottom int
}

// SnakesLadders ...
type SnakesLadders struct {
	board   [100]int
	snakes  []snake
	ladders []ladder
}

var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

// New ...
func New(numOfSnakes, numOfLadders int) SnakesLadders {
	var snakesLadders SnakesLadders
	for i := 1; i <= 100; i++ {
		snakesLadders.board[i-1] = i
	}
	for k := 0; k < numOfSnakes; k++ {
		snakesLadders.snakes = append(snakesLadders.snakes, generateSnake())
	}

	for j := 0; j < numOfLadders; j++ {
		snakesLadders.ladders = append(snakesLadders.ladders, generateLadder())
	}

	return snakesLadders
}

func generateSnake() snake {
	maxval := 99
	tail := randGenerator.Intn(maxval)
	head := randGenerator.Intn(maxval-tail+1) + tail
	return snake{
		head: head,
		tail: tail,
	}
}

func generateLadder() ladder {
	maxval := 99
	down := randGenerator.Intn(maxval)
	up := randGenerator.Intn(maxval-down+1) + down
	return ladder{
		top:    up,
		bottom: down,
	}
}
