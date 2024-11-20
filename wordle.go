package wordle

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

type letterStatus int

const (
	none letterStatus = iota
	absent
	present
	correct
)

type guess [wordSize]letter

type letter struct {
	char byte
	status string
}

func newWordleState(word string) wordleState {
	state := wordleState{}
	copy(state.word[:], word)

	return state
}

func main()  {
	
}