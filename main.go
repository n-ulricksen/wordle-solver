package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	clueWhite  rune = 'w'
	clueYellow rune = 'y'
	clueGreen  rune = 'g'
)

var validClues = map[rune]bool{
	clueWhite:  true,
	clueYellow: true,
	clueGreen:  true,
}

func main() {
	// TODO: embed word list file
	wordList := loadWordList("./words.txt")

	for len(wordList) > 1 {
		word, err := promptForWord(wordList)
		for err != nil {
			fmt.Printf("%s\n\n", err)
			word, err = promptForWord(wordList)
		}

		clues, err := promptForClues()
		for err != nil {
			fmt.Printf("%s\n\n", err)
			clues, err = promptForClues()
		}

		found := make([]bool, 5)
		for i, c := range clues {
			if c == clueGreen {
				found[i] = true
			}
		}

		// Check every word in the word list to see if it is a potential winner
		var newWordList []string
		for _, guess := range wordList {
			if isPotentialWinner(word, guess, clues, found) {
				newWordList = append(newWordList, guess)
			}
		}

		fmt.Println("\nPotential winners:")
		fmt.Printf("%s\n\n", newWordList)

		wordList = newWordList
	}
}

func promptForWord(wordList []string) (string, error) {
	var userInput string

	fmt.Println("Enter word guessed:")
	fmt.Scanf("%s", &userInput)

	if !isValidWord(wordList, userInput) {
		return "", errors.New("Invalid word.")
	}

	for _, ch := range userInput {
		if ch < 'a' || ch > 'z' {
			return "",
				errors.New("Word can only consist of lowercase English letters.")
		}
	}

	return userInput, nil
}

func promptForClues() (string, error) {
	var userInput string

	fmt.Println("Enter clues in the format `xxxxx` (w, y, g):")
	fmt.Scanf("%s", &userInput)

	if len(userInput) != 5 {
		return "", errors.New("Enter exactly 5 letters, one for each clue.")
	}

	for _, ch := range userInput {
		if !validClues[ch] {
			return "", errors.New("Invalid clues format. (w, y, g) only.")
		}
	}

	return userInput, nil
}

// isPotentialWinner checks the guess from the word list against the previously
// guessed word to determine, based on the given clues, if it is still a
// contender for being the correct word.
func isPotentialWinner(prevWord, guess, clues string, found []bool) bool {
	for i := 0; i < 5; i++ {
		switch rune(clues[i]) {
		case clueGreen:
			if !isGreenCorrect(prevWord, guess, i) {
				return false
			}
		case clueYellow:
			if !isYellowCorrect(prevWord, guess, i) {
				return false
			}
		case clueWhite:
			if !isWhiteCorrect(prevWord, guess, i, found) {
				return false
			}
		}
	}

	return true
}

// isGreenCorrect returns true if letters from both words at the given idx are
// the same.
func isGreenCorrect(word, dictGuess string, idx int) bool {
	return word[idx] == dictGuess[idx]
}

// isYellowCorrect returns true if the letter at word[idx] is found somewhere
// in dictGuess besides the given idx.
func isYellowCorrect(word, dictGuess string, idx int) bool {
	for i, dictCh := range dictGuess {
		if i == idx {
			continue
		}
		if dictCh == rune(word[idx]) {
			return true
		}
	}
	return false
}

// isWhiteCorrect returns true if the letter at word[idx] is NOT found in
// dictGuess (accept for already solved characters).
func isWhiteCorrect(word, dictGuess string, idx int, found []bool) bool {
	for i, dictCh := range dictGuess {
		if dictCh == rune(word[idx]) && !found[i] {
			return false
		}
	}
	return true
}

// isValidWord returns true if the given word is found in the given word list.
func isValidWord(wordList []string, word string) bool {
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}

func loadWordList(filepath string) []string {
	wordListTxt, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	wordList := strings.Split(strings.TrimSpace(string(wordListTxt)), "\n")

	return wordList
}
