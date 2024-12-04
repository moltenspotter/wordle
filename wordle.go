package wordle

import (
	"fmt"
	"slices"
)

const (
	maxGuesses = 6
	wordSize = 5
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [wordSize]byte
	// guesses holds the guesses that the user has made
	guesses [maxGuesses]guess
	// currGuess is the index of the available slot in guesses
	currGuess int
}

func newWordleState(word string) wordleState {
	ws := wordleState{}
	copy(ws.word[:], word)

	return ws
}

func (ws *wordleState) appendGuess(g guess) error {
	ws.guesses[ws.currGuess] = g
	ws.currGuess += 1
	return nil
}

type guess [wordSize]letter

func (g guess) String() string {
	output := ""
	for _,v := range g {
		output += fmt.Sprintf("%s\n", v)
	}
	return output
}

func (g guess) toByteSlice() []byte {
	var result []byte
	for _,v := range g {
		result = append(result, v.char)
	}
	return result
}

func (g *guess) updateLettersWithWord(correctWord [wordSize]byte) {
	for i,v := range g.toByteSlice() {
		if v == correctWord[i] {
			g[i].status = correct
		} else if slices.Contains(correctWord[:], v) {
			g[i].status = present
		} else {
			g[i].status = absent
		}
	}
}

type letter struct {
	char byte
	status letterStatus
}

func (l letter) String() string {
	return fmt.Sprintf("Char: %c, Status: %s", l.char, l.status)
}

type letterStatus int

const (
	none letterStatus = iota
	absent
	present
	correct
)

func (ls letterStatus) String() string {
	return statusToString(ls)
}

func statusToString(status letterStatus) string {
	switch status {
	case none:
		return "none"
	case correct:
		return "correct"
	case present:
		return "present"
	case absent:
		return "absent"
	default:
		return "unknown"
	}
}



func newLetter(char byte) letter {
	return letter{char, none}
}

func newGuess(inputGuess string) guess {
	outputGuess := guess{}
	for i,v := range inputGuess {
		outputGuess[i] = newLetter(byte(v))
	}
	return outputGuess
}


func main() {
	
}
