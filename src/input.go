package hangmanweb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func GetMenuChoice(prompt string, maxChoice int) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if num, err := strconv.Atoi(choice); err == nil && num >= 1 && num <= maxChoice {
			return choice
		}

		fmt.Printf("Invalid input. Please enter a number between 1 and %d.\n", maxChoice)
	}
}
