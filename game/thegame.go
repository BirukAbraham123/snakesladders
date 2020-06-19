package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	sl "github.com/BirukAbraham123/snakesladders"
	fg "github.com/mbndr/figlet4go"
)

func main() {
	gameController()
}

var diesRollGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func diesRole() int {
	return diesRollGenerator.Intn(6) + 1
}

// func tailOrHead() int {
// 	return diesRollGenerator.Intn(2)
// }

type player struct {
	position int
	name     string
}

func (player *player) String() string {
	return fmt.Sprintf("%v->Score:%d", player.name, player.position)
}

var whoStarts int
var reader *bufio.Reader

func gameController() {
	figletPrint("Snakes and Ladders")
	fmt.Println("\nWelcom the Snakes and Ladders Game !!!")

	fmt.Println("Starting the Game")
	numberOfSnakes := 6
	numberOfLadders := 7

	computer := player{position: 1, name: "computer"}

	gameBoard := sl.New(numberOfSnakes, numberOfLadders)

	reader = bufio.NewReader(os.Stdin)
	var human player
	fmt.Print("Enter your name please: ")
	playerName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	} else {
		human = player{position: 1, name: playerName}
	}

	fmt.Print("Started the game ...\n")
	printInfo(&human, &computer, &gameBoard)
	fmt.Print("Tail or Head write(t/h) : ")

	// firstChooser(reader)
	gameFlowHandler(&human, &computer, &gameBoard)
}

func gameFlowHandler(hum, com *player, gBoard *sl.SnakesLadders) {
	fmt.Print("Please you start rolling dies (r or enter): ")
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	}
	humanHandler(hum, com, gBoard, true)
}

func humanHandler(hum, com *player, gBoard *sl.SnakesLadders, flag bool) {
	f := func() {
		fmt.Println("Rolling ....")
		time.Sleep(2 * time.Second)
		result := diesRole()
		fmt.Printf("Dies faced %d\n", result)
		prePosition := hum.position
		hum.position = gBoard.Play(result, prePosition)
		if hum.position == 100 {
			figletPrint("You are the winner !!!")
		} else if hum.position >= prePosition {
			fmt.Println("Good getting closer ...")
			printInfo(hum, com, gBoard)
			fmt.Println("Now is my turn")
			aiHandler(hum, com, gBoard)
		} else if hum.position <= prePosition {
			fmt.Println("Opoos ... going down :)")
			printInfo(hum, com, gBoard)
			fmt.Println("Now is my turn")
			aiHandler(hum, com, gBoard)
		}
	}
	if flag {
		f()
	} else {
		fmt.Print("Go ahead .. :")
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong")
		} else {
			f()
		}
	}
}

func aiHandler(hum, com *player, gBoard *sl.SnakesLadders) {
	fmt.Println("AI rolling ...")
	time.Sleep(2 * time.Second)
	result := diesRole()
	fmt.Printf("Dies faced %d\n", result)
	preCom := com.position
	com.position = gBoard.Play(result, preCom)
	if com.position == 100 {
		figletPrint("You lose GAME OVER !!!")
	} else if com.position >= preCom {
		fmt.Println("I am going up ... look out")
		printInfo(hum, com, gBoard)
		fmt.Println("Now is your turn")
		humanHandler(hum, com, gBoard, false)
	} else if com.position <= preCom {
		fmt.Println("Opoos ... sh*t :(")
		printInfo(hum, com, gBoard)
		fmt.Println("Now is your turn")
		humanHandler(hum, com, gBoard, false)
	}
}

func printInfo(human, com *player, skLad *sl.SnakesLadders) {
	fmt.Printf("%v\n", human.String())
	fmt.Printf("%v\n", com.String())
	fmt.Printf("%v\n", skLad.String())
}

// func firstChooser(reader *bufio.Reader) {
// 	torh, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Something went wrong")
// 	} else {
// 		torh = strings.ToLower(strings.TrimSpace(torh))
// 		if torh == "t" || torh == "h" {
// 			whoStarts = tailOrHead()
// 		} else {
// 			fmt.Print("Please enter t or h: ")
// 			firstChooser(reader)
// 		}
// 	}
// }

func figletPrint(str string) {
	ascii := fg.NewAsciiRender()
	renderStr, _ := ascii.Render(str)
	fmt.Print(renderStr)
}
