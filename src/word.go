package hangmanweb

import (
	"math/rand"
	"strings"
	"time"
)



func NewWord(words []string) *Word {
	rand.Seed(time.Now().UnixNano())
	selectedWord := words[rand.Intn(len(words))]
	return &Word{
		fullWord:      selectedWord,
		revealedWord:  make([]rune, len(selectedWord)),
		revealedCount: 0,
	}
}

func (w *Word) RevealRandomLetters(count int) {
	for i := 0; i < count; i++ {
		if w.revealedCount >= len(w.fullWord) {
			return
		}
		idx := rand.Intn(len(w.fullWord))
		if w.revealedWord[idx] == 0 {
			w.revealedWord[idx] = rune(w.fullWord[idx])
			w.revealedCount++
		} else {
			i--
		}
	}
}

func (w *Word) RevealLetter(letter rune) bool {
	revealed := false
	for i, char := range w.fullWord {
		if char == letter && w.revealedWord[i] == 0 {
			w.revealedWord[i] = letter
			w.revealedCount++
			revealed = true
		}
	}
	return revealed
}

func (w *Word) Guess(word string) bool {
	return strings.EqualFold(word, w.fullWord)
}

func (w *Word) RevealAllLetters() {
	for i, char := range w.fullWord {
		if w.revealedWord[i] == 0 {
			w.revealedWord[i] = char
			w.revealedCount++
		}
	}
}

func (w *Word) IsFullyRevealed() bool {
	return w.revealedCount == len(w.fullWord)
}

func (w *Word) GetDisplayWord() string {
	display := make([]rune, len(w.fullWord))
	for i, char := range w.revealedWord {
		if char == 0 {
			display[i] = '_'
		} else {
			display[i] = char
		}
	}
	return string(display)
}

func (w *Word) GetFullWord() string {
	return w.fullWord
}