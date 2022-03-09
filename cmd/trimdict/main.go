package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Finds all 5 letter words from the `words_alpha.txt` word list, and saves
// a new file containing these 5 letter words.
func main() {
	wordList := loadWordList("./words_alpha.txt")

	var trimmed bytes.Buffer
	for _, word := range wordList {
		if len(word) == 5 {
			trimmed.WriteString(word)
			trimmed.WriteByte('\n')
		}
	}

	ioutil.WriteFile("./words.txt", trimmed.Bytes(), 0644)
}

func loadWordList(filepath string) []string {
	wordListTxt, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	wordList := strings.Split(string(wordListTxt), "\r\n")

	return wordList
}
