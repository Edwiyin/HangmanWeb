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

func (w *Word) GetFullWord() string {
var revealedWord strings.Builder
	for _, char := range w.revealedWord {
		if char == 0 {
			revealedWord.WriteRune('_')
		} else {
			revealedWord.WriteRune(char)
		}
	}
	return revealedWord.String()
}


func (w *Word) IsLetterRevealed(letter rune) bool {
    for i, char := range w.fullWord {
        if char == letter && w.revealedWord[i] == letter {
            return true
        }
    }
    return false
}

func (w *Word) GetDisplayWord() string {
    var displayWord strings.Builder
	for _, char := range w.revealedWord {
		if char == 0 {
			displayWord.WriteRune('_')
		} else {
			displayWord.WriteRune(char)
		}
	}
	return displayWord.String()
}

func (w *Word) IsFullyRevealed() bool {
    return w.revealedCount == len(w.fullWord)
}

func (w *Word) GetRevealedCount() int {
    count := 0
    for _, char := range w.revealedWord {
        if char != 0 {
            count++
        }
    }
    return count
}

func (w *Word) RevealLetter(letter rune) bool {
    found := false
    letter = toLower(letter)
    
    for i, char := range w.fullWord {
        if toLower(char) == letter && w.revealedWord[i] == 0 {
            w.revealedWord[i] = char
            w.revealedCount++
            found = true
        }
    }
    
    return found
}


func toLower(r rune) rune {
    if r >= 'A' && r <= 'Z' {
        return r + ('a' - 'A')
    }
    return r
}