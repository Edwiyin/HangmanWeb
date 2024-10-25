package hangmanweb

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func CreateNewGame(playerName, difficulty string) *Game {
    diff := Difficulty(strings.ToLower(difficulty))
    
    words, err := GetWordsByDiff(string(diff))
    if err != nil {
        log.Printf("Error loading words: %v, using default word list", err)
        words = []string{"ordinateur", "programmation", "interface", "application"}
    }
    
    selectedWord := selectRandomWord(words)
    
    word := &Word{
        fullWord:      selectedWord,
        revealedWord:  make([]rune, len(selectedWord)),
        revealedCount: 0,
    }
    
    game := &Game{
        PlayerName:     playerName,
        Word:          word,
        GuessedLetters: make(map[rune]string),
        MaxTries:      DifficultyConfig[diff].MaxTries,
        Diff:          diff,
        GuessedWord:   make([]string, len(selectedWord)),
    }
    
    for i := range game.GuessedWord {
        game.GuessedWord[i] = "_"
    }
   
    if diff == "easy" {
        word.RevealRandomLetters(2)

    } else if diff == "medium" {
        word.RevealRandomLetters(1)
    }
    
    for i, r := range word.revealedWord {
        if r != 0 {
            game.GuessedWord[i] = string(r)
            game.GuessedLetters[r] = string(r)
        }
    }
    
    return game
}

func GetRandomWord() {
	panic("unimplemented")
}

func selectRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}
	return words[rand.Intn(len(words))]
}

func loadWordsFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", filepath, err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filepath, err)
	}

	return words, nil
}


func GetWordsByDiff(difficulty string) ([]string, error) {
	
	rand.Seed(time.Now().UnixNano())

	difficulty = strings.ToLower(difficulty)

	filepath, exists := DifficultyFiles[difficulty]
	if !exists {
		return nil, fmt.Errorf("invalid difficulty level: %s. Choose 'easy', 'medium', or 'hard'", difficulty)
	}

	words, err := loadWordsFromFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("no words found in file for difficulty level: %s", difficulty)
	}

	return words, nil
}


func GetPlayerGuess(guessedLetters map[rune]bool) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a letter or guess the whole word: ")
		input, _ := reader.ReadString('\n')
		guess := strings.TrimSpace(strings.ToLower(input))

		if len(guess) == 1 {
			letter := rune(guess[0])
			if guessedLetters != nil && guessedLetters[letter] {
				fmt.Println("You've already guessed that letter. Try again.")
				continue
			}
			return guess
		} else if len(guess) > 1 {
			return guess
		}

		fmt.Println("Invalid input. Please enter a single letter or the whole word.")
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
		MaxTries: settings.MaxTries,
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
		g.MaxTries--
		fmt.Println(("Incorrect guess!"))
	}
}

func (g *Game) ProcessWordGuess(word string) {
	if g.Word.Guess(word) {
		g.Word.RevealAllLetters()
		fmt.Println(("Correct word guess!"))
	} else {
		g.MaxTries -= 2
		fmt.Println(("Incorrect word guess! You lose 2 tries."))
	}
}

func (g *Game) IsGameOver() bool {
	return g.MaxTries <= 0 || g.Word.IsFullyRevealed()
}

func (g *Game) DisplayGameState() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Word: %s\n",(g.Word.GetDisplayWord()))
	fmt.Printf("Remaining tries: %s\n", (fmt.Sprintf("%d", g.MaxTries)))
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
