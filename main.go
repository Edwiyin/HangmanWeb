package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	hangmanweb "hangmanweb/src"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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
	
	data := hangmanweb.PageData{
		Title:       "Hangman Web - Game",
		Game:        *game,
		Message:     fmt.Sprintf("Welcome %s! Game started on %s difficulty", playerName, difficulty),
		MessageType: "info",
		Scores:      []hangmanweb.Score{},
	}

	if err := tmpl.ExecuteTemplate(w, "game.html", data); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/game", handleGame)

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}