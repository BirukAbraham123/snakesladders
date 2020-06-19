package snakesladders

import (
	"fmt"
	"math/rand"
	"time"
)

const maxVal = 99

type snake struct {
	head int
	tail int
}

type ladder struct {
	top    int
	bottom int
}

// SnakesLadders a struct representation of the snakes and ladders
// board game
type SnakesLadders struct {
	board   [100]int
	snakes  map[int]int // map[snake.head]snake.tail
	ladders map[int]int // map[ladder.bottom]ladder.top
}

// Play a method which return the next position of the player
// after the dies is roled given the die face value and the previous
// position of the player (position of the player is assumed to be
// between 1 upto 100)
func (snakeladder *SnakesLadders) Play(diesVal, playerPosition int) int {
	var nextPosition int
	// compute the new position of the player ignoring
	// the ladders and the snakes as well as the last
	// exact move to the final 99 cell
	nextPosition = (playerPosition - 1) + diesVal

	if nextPosition > maxVal {
		nextPosition = maxVal + (maxVal - nextPosition)
	} else {
		// check if player is on the head of a snake
		for head, tail := range snakeladder.snakes {
			if nextPosition == head {
				nextPosition = tail
			}
		}

		// check if the player is on the bottom of a ladder
		for bottom, top := range snakeladder.ladders {
			if nextPosition == bottom {
				nextPosition = top
			}
		}
	}

	return snakeladder.board[nextPosition]
}

func (snakeladder *SnakesLadders) String() string {
	p := snakeladder.board
	snakesList := make([]string, len(snakeladder.snakes))
	for shead, stail := range snakeladder.snakes {
		snakesList = append(snakesList, fmt.Sprintf("snake tail : %d ---> head : %d", stail, shead))
	}
	laddersList := make([]string, len(snakeladder.ladders))
	for lbottom, ltop := range snakeladder.ladders {
		laddersList = append(laddersList, fmt.Sprintf("ladder bottom : %d ----> top : %d", lbottom, ltop))
	}
	boardStr := "_________________________________________\n"
	for i := 99; i >= 0; i = i - 10 {
		temp := ""
		if len(laddersList) != 0 {
			temp = laddersList[len(laddersList)-1]
			laddersList = laddersList[:len(laddersList)-1] // removed a ladder item
		}
		if len(snakesList) != 0 {
			temp += fmt.Sprintf("\t%v", snakesList[len(snakesList)-1])
			snakesList = snakesList[:len(snakesList)-1] // removed a snake item
		}
		if int((i-9)/10)%2 != 0 {
			if temp != "" {
				boardStr += fmt.Sprintf("|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|\t%v\n", p[i], p[i-1], p[i-2], p[i-3], p[i-4], p[i-5], p[i-6], p[i-7], p[i-8], p[i-9], temp)
			} else {
				boardStr += fmt.Sprintf("|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|\n", p[i], p[i-1], p[i-2], p[i-3], p[i-4], p[i-5], p[i-6], p[i-7], p[i-8], p[i-9])
			}
		} else {
			if temp != "" {
				boardStr += fmt.Sprintf("|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|\t%v\n", p[i-9], p[i-8], p[i-7], p[i-6], p[i-5], p[i-4], p[i-3], p[i-2], p[i-1], p[i], temp)
			} else {
				boardStr += fmt.Sprintf("|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|%3d|\n", p[i-9], p[i-8], p[i-7], p[i-6], p[i-5], p[i-4], p[i-3], p[i-2], p[i-1], p[i])
			}
		}
		boardStr += fmt.Sprintln("-----------------------------------------")
	}
	return boardStr
}

var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

// New a function which creates a Snakes and Ladders board instance
// given the number of snakes and number of ladders as integer values
func New(numOfSnakes, numOfLadders int) SnakesLadders {
	var snakesLadders SnakesLadders
	for i := 1; i <= 100; i++ {
		snakesLadders.board[i-1] = i
	}

	snakesLadders.snakes = make(map[int]int, numOfSnakes)
	for k := 0; k < numOfSnakes; k++ {
		generateSnake(&snakesLadders, k+1)
	}

	snakesLadders.ladders = make(map[int]int, numOfLadders)
	for j := 0; j < numOfLadders; j++ {
		generateLadder(&snakesLadders)
	}

	return snakesLadders
}

func generateSnake(snakeladder *SnakesLadders, index int) {
	// tail := randGenerator.Intn(maxVal - 1)
	// potentialHead := randGenerator.Intn((maxVal-tail)+1) + tail
	tail := randGenerator.Intn(int(maxVal/index) - 1)
	potentialHead := randGenerator.Intn((maxVal-tail)+1) + tail

	if potentialHead == maxVal {
		// the head of the snake will be the winning position
		potentialHead = potentialHead - 1
	}
	if _, ok := snakeladder.snakes[potentialHead]; !ok {
		if potentialHead <= tail {
			generateSnake(snakeladder, index)
		} else {
			snakeladder.snakes[potentialHead] = tail
		}
	} else {
		generateSnake(snakeladder, index)
	}
}

func generateLadder(snakeladder *SnakesLadders) {
	potentialDown := randGenerator.Intn(maxVal - 1)
	if potentialDown == 0 {
		potentialDown = potentialDown + 1
	}
	if _, ok := snakeladder.ladders[potentialDown]; !ok {
		up := randGenerator.Intn((maxVal-potentialDown)+1) + potentialDown
		if up == potentialDown {
			generateLadder(snakeladder)
		} else {
			snakeladder.ladders[potentialDown] = up
		}
	} else {
		generateLadder(snakeladder)
	}
}
