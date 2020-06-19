package snakesladders

import (
	"fmt"
	"math/rand"
	"testing"
)

var numberOfSnakes = 6
var numberOfLadders = 7

func TestNew(t *testing.T) {
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

	for head, tail := range snakeladder.snakes {
		if head <= tail {
			t.Error("Expected snake head to be always greater than tail position")
		}
	}

	for bottom, top := range snakeladder.ladders {
		if top <= bottom {
			t.Error("Expected ladder top to be always greater than bottom position")
		}
	}
}

func TestPlay(t *testing.T) {
	// To make the snake and ladder map deterministic
	// 	Snake Map [head]tail :  map[28:27 56:25 82:14 85:6 88:26 96:11]
	// Ladder Map [bottom]up :  map[8:46 20:66 23:47 45:88 46:88 59:68 76:83]
	randGenerator = rand.New(rand.NewSource(0))
	snakeladder := New(numberOfSnakes, numberOfLadders)

	// if player was on position 1 and dies roll is 6, player position should
	// be 7
	playerPositionOne := 7
	resultPositionOne := snakeladder.Play(6, 1)
	if resultPositionOne != playerPositionOne {
		t.Errorf("snakeladder.Play(6, 1) : got %d should be %d", resultPositionOne, playerPositionOne)
	}

	// if player was on position 52 and dies roll is 5, player position should
	// be 26, player eaten by the snake and moves back to position 26
	playerPositionTwo := 26
	resultPositionTwo := snakeladder.Play(5, 52)
	if resultPositionTwo != playerPositionTwo {
		t.Errorf("snakeladder.Play(5, 52) : got %d should be %d", resultPositionTwo, playerPositionTwo)
	}

	// if player was on position 4 and dies roll is 5, player position should
	// be on 89 because the player moves by the first ladder to position 47 then
	// moves to position 89 by another ladder on position 47
	playerPositionThr := 89
	resultPositionThr := snakeladder.Play(5, 4)
	if resultPositionThr != playerPositionThr {
		t.Errorf("snakeladder.Play(5, 4) : got %d should be %d", resultPositionThr, playerPositionThr)
	}

	// if player was on position 98 and dies role is 5, player moves to winning
	// position which is 100 and moves back by three steps to position 97
	playerPositionFour := 97
	resultPositionFour := snakeladder.Play(5, 98)
	if resultPositionFour != playerPositionFour {
		t.Errorf("snakeladder.Play(5, 98) : got %d should be %d", resultPositionFour, playerPositionFour)
	}

	// if player was on position 98 and dies role is 2, player moves to winning
	// position which is 100
	playerPositionFive := 100
	resultPositionFive := snakeladder.Play(2, 98)
	if resultPositionFive != playerPositionFive {
		t.Errorf("snakeladder.Play(2, 98) : got %d should be %d", resultPositionFive, playerPositionFive)
	}
}

func TestString(t *testing.T) {
	snakeladder := New(numberOfSnakes, numberOfLadders)
	fmt.Println(snakeladder.String())
}
