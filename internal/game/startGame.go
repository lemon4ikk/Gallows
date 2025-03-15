package game

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	lettersFile    = "list_of_letters.txt"
	dictionaryFile = "dictionary.txt"
)

func Start() {
	var dictionary []string

	dictionaryDescriptor, err := os.Open(dictionaryFile)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %s", dictionaryFile)
	}

	infoDictionary, _ := dictionaryDescriptor.Stat()

	if infoDictionary.Size() == 0 {
		log.Fatalf("Файл %s пустой", dictionaryFile)
	}

	scanWord := bufio.NewScanner(dictionaryDescriptor)
	for scanWord.Scan() {
		dictionary = append(dictionary, scanWord.Text())
	}

	wordIndex := rand.Intn(len(dictionary))
	dictionaryDescriptor.Close()

	word := dictionary[wordIndex]

	var underscore []string
	var expectedWord string

	for range word {
		underscore = append(underscore, "_")
	}

	expectedWord = strings.Join(underscore, " ")

	secretWord := make([]rune, utf8.RuneCountInString(expectedWord))

	for i, char := range expectedWord {
		secretWord[i] = char
	}

	expectedWord = ""
	for _, char := range word {
		expectedWord += string(char) + " "
	}
	secretWord = append(secretWord, ' ')

	GameLogic(lettersFile, secretWord, expectedWord, word)
}
