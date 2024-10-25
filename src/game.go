package hangmanweb

import (
	"fmt"
	"strings"
)

func CreateNewGame(playerName, difficulty string) *Game {
	wordStr := "example"
	guessedWord := make([]string, len(wordStr))
	for i := range guessedWord {
		guessedWord[i] = "_"
	}

	word := &Word{Value: wordStr} 

	return &Game{
		GuessedWord:    guessedWord,
		PlayerName:     playerName,
		Word:           word,
		GuessedLetters: make(map[rune]string),
		RemainingTries: 10,
		Diff:     Difficulty(difficulty),
	}
}

func mapKeysToString(m map[rune]string) []string {
	var array []string = []string{}
	for _, v := range m {
		array = append(array, string(v))
	}
	return array
}

func NewGame(words []string, difficulty Difficulty) *Game {
	settings := DifficultyConfig[difficulty]
	return &Game{
		Word:           NewWord(words),
		GuessedLetters: make(map[rune]string),
		RemainingTries: settings.MaxTries,
		Diff:     difficulty,
	}
}

func (g *Game) Play() {
	settings := DifficultyConfig[g.Diff]
	g.Word.RevealRandomLetters(settings.InitialReveals)

	for !g.IsGameOver() {
		g.DisplayGameState()
		guessedLettersBool := make(map[rune]bool)
		for k := range g.GuessedLetters {
			guessedLettersBool[k] = true
		}
		guess := GetPlayerGuess(guessedLettersBool)

		if len(guess) == 1 {
			g.ProcessLetterGuess(rune(guess[0]))
		} else {
			g.ProcessWordGuess(guess)
		}
	}
}

func (g *Game) ProcessLetterGuess(letter rune) {
	if g.Word.RevealLetter(letter) {
		g.GuessedLetters[letter] = (string(letter))
		fmt.Println(("Correct guess!"))
	} else {
		g.GuessedLetters[letter] = (string(letter))
		g.RemainingTries--
		fmt.Println(("Incorrect guess!"))
	}
}

func (g *Game) ProcessWordGuess(word string) {
	if g.Word.Guess(word) {
		g.Word.RevealAllLetters()
		fmt.Println(("Correct word guess!"))
	} else {
		g.RemainingTries -= 2
		fmt.Println(("Incorrect word guess! You lose 2 tries."))
	}
}

func (g *Game) IsGameOver() bool {
	return g.RemainingTries <= 0 || g.Word.IsFullyRevealed()
}

func (g *Game) DisplayGameState() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Word: %s\n",(g.Word.GetDisplayWord()))
	fmt.Printf("Remaining tries: %s\n", (fmt.Sprintf("%d", g.RemainingTries)))
	fmt.Printf("Guessed letters: ")
	array := mapKeysToString(g.GuessedLetters)
	for _, v := range array {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
	fmt.Println(strings.Repeat("-", 40))
}

func (g *Game) CheckGuessedLetters(input string) bool {
	for _, letter := range g.GuessedLetters {
		if letter == input {
			return true
		}
	}
	return false
}
