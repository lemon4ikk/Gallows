package main

import (
	"bufio"
	"fmt"
	"gallows/game"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var userMsg string
	const (
		startCommand = "Начать новую игру"
		exitCommand  = "Выйти"
	)

	fmt.Println("Для начала игры введите: " + startCommand)
	fmt.Println("Для выхода из игры введите: " + exitCommand)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userMsg = scanner.Text()

	if userMsg == exitCommand {
		os.Exit(0)
	}

	if userMsg == startCommand {
		Start()
	}
}

func Start() {
	const (
		lettersFile    = "list_of_letters.txt"
		dictionaryFile = "dictionary.txt"
	)
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

	game.GameLogic(lettersFile, secretWord, expectedWord, word)
}
