package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type QuestionCategory int

const (
	Pop QuestionCategory = iota
	Science
	Sports
	Rock
)

func (q QuestionCategory) String() string {
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

type QuestionDeck []string

type Game struct {
	players            []Player
	currentPlayerIndex int

	questionDecks map[QuestionCategory]QuestionDeck

	isGettingOutOfPenaltyBox bool
}

func NewGame() *Game {
	game := &Game{}

	game.questionDecks = map[QuestionCategory]QuestionDeck{Pop: {}, Science: {}, Sports: {}, Rock: {}}
	for key, value := range game.questionDecks {
		for i := 0; i < 50; i++ {
			value = append(value, fmt.Sprintf("%v Question %d\n", key, i))
		}
		game.questionDecks[key] = value
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
	questionCategory := getPlaceQuestionCategory(me.getCurrentPlayer().place)

	fmt.Printf("%s's new location is %d\n", me.getCurrentPlayer().name, me.getCurrentPlayer().place)
	fmt.Printf("The category is %s\n", questionCategory)
	me.askQuestion(questionCategory)
}

func (me *Game) askQuestion(questionCategory QuestionCategory) {
	questions := me.questionDecks[questionCategory]
	question := questions[0]
	me.questionDecks[questionCategory] = questions[1:]
	fmt.Print(question)
}

func getPlaceQuestionCategory(place int) QuestionCategory {
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
