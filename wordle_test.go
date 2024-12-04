package wordle

import (
	"slices"
	"testing"
	words "github.com/moltenspotter/wordle/words"
)

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

func TestNewGuess(t *testing.T) {
	guessString := "AUDIO"
	generatedGuess := newGuess(guessString)

	t.Logf("Guess String: %s\n", guessString)
	t.Logf("Generated guess:\n%s\n", generatedGuess)
	for i,v := range generatedGuess {
		if v != newLetter(guessString[i]) {
			t.Errorf("Generated guess letter %d should be %s", i, v)
		}
	}
}

func TestUpdateLettersWithWord (t *testing.T) {
	correctWord := [wordSize]byte{}
	copy(correctWord[:], "YIELD")

	guessWord := "BILLY"

	generatedGuess := newGuess(guessWord)
	generatedGuess.updateLettersWithWord(correctWord)

	correctStatusArr := [wordSize]letterStatus{absent, correct, present, correct, present}
	for i,l := range generatedGuess {
		if l.status != correctStatusArr[i] {
			t.Errorf("Letter %d (%c) expected status: %s, actually %s", l.char, i, correctStatusArr[i], l.status)
		}
	}
}

func TestToByteSlice (t *testing.T) {
	g := newGuess("ABCDE")
	bs := g.toByteSlice()

	expected := []byte{byte('A'), byte('B'), byte('C'), byte('D'), byte('E')}

	if !(slices.Equal(bs, expected)) {
		t.Errorf("Expected %c, got %c", expected, bs)
	}
}

func  TestAppendGuess(t *testing.T) {
	ws := newWordleState("YIELD")

	guessWord := "HILLY"
	g := newGuess(guessWord)
	err := ws.appendGuess(g)
	if ws.guesses[0] != g {
		t.Errorf("Did not append guess to guesses")
	} else if len(ws.guesses) >= maxGuesses && err != nil {
		t.Errorf("Should have errored when adding too many guesses")
	} else if len(g) != wordSize && err != nil {
		t.Errorf("Should have errored due to guess being wrong size")
	} else if !(words.IsWord(guessWord)) && err != nil {
		t.Errorf("Should have errored due to guess not being a valid word")
	}


}

