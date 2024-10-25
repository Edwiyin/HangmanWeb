package hangmanweb

import (
	"time"
)

type PageData struct {
	Title          string
	Game           Game
	Message        string
	MessageType    string
	Scores         []Score
	CurrentSession string
}

type Game struct {
	 GuessedWord    []string
    PlayerName     string
    GuessedLetters map[rune]string
    Diff          Difficulty
    MaxTries      int
    Word          *Word
}

type Word struct {
	fullWord      string
	revealedWord  []rune
	revealedCount int
	Value         string
}

type Score struct {
	PlayerName string    `json:"playerName"`
	Score      int       `json:"score"`
	Word       string    `json:"word"`
	Difficulty string    `json:"difficulty"`
	Date       time.Time `json:"date"`
}
