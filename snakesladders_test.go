package snakesladders

import (
	"testing"
)

func TestNew(t *testing.T) {
	numberOfSnakes := 6
	numberOfLadders := 7
	snakeladder := New(numberOfSnakes, numberOfLadders)
	expectedBoard := [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

	if len(snakeladder.snakes) != numberOfSnakes {
		t.Errorf("Number of snakes expected %d got %d", numberOfSnakes, len(snakeladder.snakes))
	}
	if len(snakeladder.ladders) != numberOfLadders {
		t.Errorf("Number of ladders expected %d got %d", numberOfLadders, len(snakeladder.ladders))
	}
	if expectedBoard != snakeladder.board {
		t.Error("Expected board does has conflict : got ", snakeladder.board)
	}

	for _, snake := range snakeladder.snakes {
		if snake.head <= snake.tail {
			t.Error("Expected snake head to be always greater than tail position")
		}
	}

	for _, ladder := range snakeladder.ladders {
		if ladder.top <= ladder.bottom {
			t.Error("Expected ladder top to be always greater than bottom position")
		}
	}
}
