package hangmanweb

type Difficulty string

const (
	Easy   Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard   Difficulty = "Hard"
)

type DifficultySettings struct {
	InitialReveals int
	MaxTries       int
}

var DifficultyConfig = map[Difficulty]DifficultySettings{
	Easy:   {InitialReveals: 2, MaxTries: 10},
	Medium: {InitialReveals: 1, MaxTries: 10},
	Hard:   {InitialReveals: 0, MaxTries: 10},
}