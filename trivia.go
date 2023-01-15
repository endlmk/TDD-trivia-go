package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type QuizCategory int

const (
	Pop QuizCategory = iota
	Science
	Sports
	Rock
)

func (q QuizCategory) String() string {
	switch q {
	case Pop:
		return "Pop"
	case Science:
		return "Science"
	case Sports:
		return "Sports"
	case Rock:
		return "Rock"
	default:
		return "unknown value"
	}
}

type Player struct {
	name         string
	place        int
	purse        int
	inPenaltyBox bool
}

type Game struct {
	players            []Player
	currentPlayerIndex int

	popQuestions     []string
	scienceQuestions []string
	sportsQuestions  []string
	rockQuestions    []string

	isGettingOutOfPenaltyBox bool
}

func NewGame() *Game {
	game := &Game{}

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
	me.players = append(me.players, Player{name: playerName})

	fmt.Printf("%s was added\n", playerName)
	fmt.Printf("They are player number %d\n", me.howManyPlayers())

	return true
}

func (me *Game) Roll(roll int) {
	fmt.Printf("%s is the current player\n", me.getCurrentPlayer().name)
	fmt.Printf("They have rolled a %d\n", roll)

	if me.getCurrentPlayer().inPenaltyBox {
		if roll%2 != 0 {
			me.isGettingOutOfPenaltyBox = true
			fmt.Printf("%s is getting out of the penalty box\n", me.getCurrentPlayer().name)
		} else {
			me.isGettingOutOfPenaltyBox = false
			fmt.Printf("%s is not getting out of the penalty box\n", me.getCurrentPlayer().name)
			return
		}
	}
	me.getCurrentPlayer().place = (me.getCurrentPlayer().place + roll) % 12
	quizCategory := getPlaceQuizCategory(me.getCurrentPlayer().place)

	fmt.Printf("%s's new location is %d\n", me.getCurrentPlayer().name, me.getCurrentPlayer().place)
	fmt.Printf("The category is %s\n", quizCategory)
	me.askQuestion(quizCategory)
}

func (me *Game) askQuestion(quizCategory QuizCategory) {
	var questions *[]string
	switch quizCategory {
	case Pop:
		questions = &me.popQuestions
	case Science:
		questions = &me.scienceQuestions
	case Sports:
		questions = &me.sportsQuestions
	case Rock:
		questions = &me.rockQuestions
	default:
		return
	}
	question := (*questions)[0]
	*questions = (*questions)[1:]
	fmt.Print(question)
}

func getPlaceQuizCategory(place int) QuizCategory {
	switch place {
	case 0, 4, 8:
		return Pop
	case 1, 5, 9:
		return Science
	case 2, 6, 10:
		return Sports
	default:
		return Rock
	}
}

func (me *Game) WasCorrectlyAnswered() bool {
	if !me.getCurrentPlayer().inPenaltyBox || me.isGettingOutOfPenaltyBox {
		fmt.Println("Answer was correct!!!!")
		me.getCurrentPlayer().purse += 1
		fmt.Printf("%s now has %d Gold Coins.\n", me.getCurrentPlayer().name, me.getCurrentPlayer().purse)

		return me.didPlayerWin()
	} else {
		return false
	}
}

func (me *Game) didPlayerWin() bool {
	return me.getCurrentPlayer().purse == 6
}

func (me *Game) WrongAnswer() {
	fmt.Println("Question was incorrectly answered")
	fmt.Printf("%s was sent to the penalty box\n", me.getCurrentPlayer().name)
	me.getCurrentPlayer().inPenaltyBox = true
}

func (me *Game) nextTurn() {
	me.currentPlayerIndex = (me.currentPlayerIndex + 1) % me.howManyPlayers()
}

func (me *Game) getCurrentPlayer() *Player {
	return &me.players[me.currentPlayerIndex]
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
		game.nextTurn()
	}
}
