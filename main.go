package main

import (
	"encoding/gob"
	"fmt"
	hangmanweb "hangmanweb/src"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var (
    tmpl    *template.Template
    store   = sessions.NewCookieStore([]byte("secret-key-123"))
)

func init() {
    
    gob.Register(&hangmanweb.Game{})
    gob.Register(&hangmanweb.Word{})
    tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func handleGuess(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/game", http.StatusSeeOther)
        return
    }

    session, err := store.Get(r, "game-session")
    if err != nil {
        log.Printf("Session error: %v", err)
        http.Error(w, "Session error", http.StatusInternalServerError)
        return
    }

    if err := r.ParseForm(); err != nil {
        log.Printf("Form parsing error: %v", err)
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    gameInterface := session.Values["game"]
    if gameInterface == nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    game := gameInterface.(*hangmanweb.Game)
    guess := r.PostForm.Get("guess")

    var message string
    var messageType string

    if game.MaxTries > 0 && !isWordComplete(game.GuessedWord) {
        if len(guess) > 0 {
            letter := rune(strings.ToLower(guess)[0])
            
            if _, exists := game.GuessedLetters[letter]; !exists {
          
                fullWord := game.Word.GetFullWord()
              
                letterFound := false
                for i, char := range fullWord {
                    if strings.ToLower(string(char)) == string(letter) {
                        game.GuessedWord[i] = string(char)
                        letterFound = true
                    }
                }

                game.GuessedLetters[letter] = string(letter)

                if letterFound {
                    message = "Correct guess!"
                    messageType = "info"
                  
                    if isWordComplete(game.GuessedWord) {
                        message = "Congratulations! You've won!"
                        messageType = "success"
                    }
                } else {
                    message = "Incorrect guess!"
                    messageType = "error"
                    game.MaxTries--
                    
                    if game.MaxTries <= 0 {
                        message = fmt.Sprintf("Game Over! The word was: %s", fullWord)
                        messageType = "error"
                    }
                }
            } else {
                message = "Letter already guessed!"
                messageType = "warning"
            }
        } else {
            message = "Please enter a letter!"
            messageType = "warning"
        }
    }

    if game.MaxTries <= 0 {
        message = fmt.Sprintf("Game Over! The word was: %s", game.Word.GetFullWord())
        messageType = "error"
    } else if isWordComplete(game.GuessedWord) {
        message = "Congratulations! You've won!"
        messageType = "success"
    }

    session.Values["game"] = game
    if err := session.Save(r, w); err != nil {
        log.Printf("Session save error: %v", err)
        http.Error(w, "Session save error", http.StatusInternalServerError)
        return
    }

    data := hangmanweb.PageData{
        Title:       "Hangman Web - Game",
        Game:        *game,
        Message:     message,
        MessageType: messageType,
    }

    if err := tmpl.ExecuteTemplate(w, "game.html", data); err != nil {
        log.Printf("Template execution error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}


func isWordComplete(guessedWord []string) bool {
    for _, letter := range guessedWord {
        if letter == "_" {
            return false
        }
    }
    return true
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    session, _ := store.Get(r, "game-session")
    session.Values["game"] = nil
    session.Save(r, w)
    
    data := hangmanweb.PageData{
        Title: "Welcome to Hangman Web",
    }
    if err := tmpl.ExecuteTemplate(w, "home.html", data); err != nil {
        log.Printf("Template execution error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func handleGame(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    if err := r.ParseForm(); err != nil {
        log.Printf("Form parsing error: %v", err)
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    playerName := r.PostForm.Get("playerName")
    difficulty := r.PostForm.Get("difficulty")

    if playerName == "" || difficulty == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    game := hangmanweb.CreateNewGame(playerName, difficulty)

    session, _ := store.Get(r, "game-session")
    session.Values["game"] = game
    session.Save(r, w)

    data := hangmanweb.PageData{
        Title:       "Hangman Web - Game",
        Game:        *game,
        Message:     fmt.Sprintf("Welcome %s! Game started on %s difficulty", playerName, difficulty),
        MessageType: "info",
    }

    if err := tmpl.ExecuteTemplate(w, "game.html", data); err != nil {
        log.Printf("Template execution error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/game", handleGame)
	http.HandleFunc("/guess", handleGuess)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}