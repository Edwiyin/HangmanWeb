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
		word:           word,
		guessedLetters: make(map[rune]string),
		remainingTries: 10,
		difficulty:     Difficulty(difficulty),
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
		word:           NewWord(words),
		guessedLetters: make(map[rune]string),
		remainingTries: settings.MaxTries,
		difficulty:     difficulty,
	}
}

func (g *Game) Play() {
	fmt.Printf("\nStarting a new game on %s difficulty!\n", g.difficulty)
	settings := DifficultyConfig[g.difficulty]
	g.word.RevealRandomLetters(settings.InitialReveals)

	for !g.IsGameOver() {
		g.DisplayGameState()
		guessedLettersBool := make(map[rune]bool)
		for k := range g.guessedLetters {
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
	if g.word.RevealLetter(letter) {
		g.guessedLetters[letter] = (string(letter))
		fmt.Println(("Correct guess!"))
	} else {
		g.guessedLetters[letter] = (string(letter))
		g.remainingTries--
		fmt.Println(("Incorrect guess!"))
	}
}

func (g *Game) ProcessWordGuess(word string) {
	if g.word.Guess(word) {
		g.word.RevealAllLetters()
		fmt.Println(("Correct word guess!"))
	} else {
		g.remainingTries -= 2
		fmt.Println(("Incorrect word guess! You lose 2 tries."))
	}
}

func (g *Game) IsGameOver() bool {
	return g.remainingTries <= 0 || g.word.IsFullyRevealed()
}

func (g *Game) DisplayGameState() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Word: %s\n",(g.word.GetDisplayWord()))
	fmt.Printf("Remaining tries: %s\n", (fmt.Sprintf("%d", g.remainingTries)))
	fmt.Printf("Guessed letters: ")
	array := mapKeysToString(g.guessedLetters)
	for _, v := range array {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
	fmt.Println(strings.Repeat("-", 40))
}

func (g *Game) CheckGuessedLetters(input string) bool {
	for _, letter := range g.guessedLetters {
		if letter == input {
			return true
		}
	}
	return false
}
