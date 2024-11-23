package dutch_dictionary

import (
	"github.com/mikelangelon/dutch-dictionary/assets"
	"gopkg.in/yaml.v3"
	"log/slog"
	"math/rand"
)

type Store struct {
	Words []*Word

	Flexibility int
}

func New() *Store {
	ws := &Store{Words: parseWords()}
	ws.Shuffle()
	ws.Flexibility = 1
	return ws
}

func (ws *Store) RandomWord() *Word {
	return ws.Words[rand.Intn(len(ws.Words))]
}

func (ws *Store) Shuffle() {
	rand.Shuffle(len(ws.Words), func(i, j int) {
		ws.Words[i], ws.Words[j] = ws.Words[j], ws.Words[i]
	})
}
func (ws *Store) WordDifficulty(difficulty int) *Word {
	for i, v := range ws.Words {
		if v.Difficulty >= difficulty-ws.Flexibility && v.Difficulty <= difficulty+ws.Flexibility {
			ws.Words = append(ws.Words[0:i], append(ws.Words[i+1:], ws.Words[i])...)
			return v
		}
	}
	return nil
}

func parseWords() []*Word {
	var words []*Word
	err := yaml.Unmarshal(assets.Nouns, &words)
	if err != nil {
		slog.Error("error unmarshalling words", "error", err)
	}
	return words
}
