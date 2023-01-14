package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	players      []string
	places       []int
	purses       []int
	inPenaltyBox []bool

	popQuestions     []string
	scienceQuestions []string
	sportsQuestions  []string
	rockQuestions    []string

	currentPlayer            int
	isGettingOutOfPenaltyBox bool
}

func NewGame() *Game {
	game := &Game{}
	for i := 0; i < 6; i++ {
		game.places = append(game.places, 0)
		game.purses = append(game.purses, 0)
		game.inPenaltyBox = append(game.inPenaltyBox, false)
	}

	for i := 0; i < 50; i++ {
		game.popQuestions = append(game.popQuestions,
			fmt.Sprintf("Pop Question %d\n", i))
		game.scienceQuestions = append(game.scienceQuestions,
			fmt.Sprintf("Science Question %d\n", i))
		game.sportsQuestions = append(game.sportsQuestions,
			fmt.Sprintf("Sports Question %d\n", i))
		game.rockQuestions = append(game.rockQuestions,
			fmt.Sprintf("Rock Question %d\n", i))
	}

	return game
}

func (me *Game) IsPlayable() bool {
	return me.howManyPlayers() >= 2
}

func (me *Game) howManyPlayers() int {
	return len(me.players)
}

func (me *Game) Add(playerName string) bool {
	me.players = append(me.players, playerName)
	me.places[me.howManyPlayers()] = 0
	me.purses[me.howManyPlayers()] = 0
	me.inPenaltyBox[me.howManyPlayers()] = false

	fmt.Printf("%s was added\n", playerName)
	fmt.Printf("They are player number %d\n", me.howManyPlayers())

	return true
}

func (me *Game) Roll(roll int) {
	fmt.Printf("%s is the current player\n", me.players[me.currentPlayer])
	fmt.Printf("They have rolled a %d\n", roll)

	if me.inPenaltyBox[me.currentPlayer] {
		if roll%2 != 0 {
			me.isGettingOutOfPenaltyBox = true
			fmt.Printf("%s is getting out of the penalty box\n", me.players[me.currentPlayer])
		} else {
			me.isGettingOutOfPenaltyBox = false
			fmt.Printf("%s is not getting out of the penalty box\n", me.players[me.currentPlayer])
			return
		}
	}
	me.places[me.currentPlayer] = (me.places[me.currentPlayer] + roll) % 12
	fmt.Printf("%s's new location is %d\n", me.players[me.currentPlayer], me.places[me.currentPlayer])
	fmt.Printf("The category is %s\n", me.currentCategory())
	me.askQuestion()
}

func (me *Game) askQuestion() {
	var questions *[]string
	switch me.currentCategory() {
	case "Pop":
		questions = &me.popQuestions
	case "Science":
		questions = &me.scienceQuestions
	case "Sports":
		questions = &me.sportsQuestions
	case "Rock":
		questions = &me.rockQuestions
	default:
		return
	}
	question := (*questions)[0]
	*questions = (*questions)[1:]
	fmt.Print(question)
}

func (me *Game) currentCategory() string {
	switch me.places[me.currentPlayer] {
	case 0, 4, 8:
		return "Pop"
	case 1, 5, 9:
		return "Science"
	case 2, 6, 10:
		return "Sports"
	default:
		return "Rock"
	}
}

func (me *Game) WasCorrectlyAnswered() bool {
	if !me.inPenaltyBox[me.currentPlayer] || me.isGettingOutOfPenaltyBox {
		fmt.Println("Answer was correct!!!!")
		me.purses[me.currentPlayer] += 1
		fmt.Printf("%s now has %d Gold Coins.\n", me.players[me.currentPlayer], me.purses[me.currentPlayer])

		return me.didPlayerWin()
	} else {
		return false
	}
}

func (me *Game) didPlayerWin() bool {
	return me.purses[me.currentPlayer] == 6
}

func (me *Game) WrongAnswer() {
	fmt.Println("Question was incorrectly answered")
	fmt.Printf("%s was sent to the penalty box\n", me.players[me.currentPlayer])
	me.inPenaltyBox[me.currentPlayer] = true
}

func main() {
	seed := time.Now().UTC().UnixNano()
	fmt.Printf("%v\n", strconv.FormatInt(seed, 10))
	gameLoop(seed)
}

func gameLoop(seed int64) {
	game := NewGame()

	game.Add("Chet")
	game.Add("Pat")
	game.Add("Sue")

	rand.Seed(seed)

	for {
		game.Roll(rand.Intn(5) + 1)

		if rand.Intn(9) == 7 {
			game.WrongAnswer()
		} else {
			isWon := game.WasCorrectlyAnswered()
			if isWon {
				break
			}
		}
		game.currentPlayer = (game.currentPlayer + 1) % game.howManyPlayers()
	}
}
