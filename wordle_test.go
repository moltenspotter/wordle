package wordle

import "testing"

func TestNewWordleState(t *testing.T) {
	word := "HELLO"
	ws := newWordleState(word)
	wordleAsString := string(ws.word[:])

	t.Log("Created wordleState:")
	t.Logf("    word: %s", wordleAsString)
	t.Logf("    guesses: %v", ws.guesses)
	t.Logf("    currGuess: %v", ws.currGuess)

	if wordleAsString != word {
		t.Errorf("word should be %s but is actually %s", word, wordleAsString)
	}
}