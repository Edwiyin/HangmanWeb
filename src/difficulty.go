package hangmanweb

type Difficulty string

const (
    Easy   Difficulty = "easy"
    Medium Difficulty = "medium"
    Hard   Difficulty = "hard"
)

type DifficultySettings struct {
    InitialReveals int
    MaxTries      int
}

var DifficultyConfig = map[Difficulty]DifficultySettings{
    Easy:   {InitialReveals: 2, MaxTries: 10},
    Medium: {InitialReveals: 1, MaxTries: 7},
    Hard:   {InitialReveals: 0, MaxTries: 5},
}

var DifficultyFiles = map[string]string{
    "easy":   "words_easy.txt",
    "medium": "words_medium.txt",
    "hard":   "words_hard.txt",
}
