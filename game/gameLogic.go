package game

import (
	"bufio"
	"fmt"
	"gallows/draw"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	maxUserErrors     = 6
	stepWord          = 2
	numSecondsWaiting = 2
)

func GameLogic(filename string, secretWord []rune, expectedWord string, word string) {
	var numUserErr int
	var flagErr bool
	var char string
	wrongLetters := []string{}
	letters := []string{}

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не удалось %s открыть файл", filename)
	}

	info, _ := inputFile.Stat()

	if info.Size() == 0 {
		log.Fatal("Файл пустой")
	}

	input := bufio.NewScanner(inputFile)
	for input.Scan() {
		letters = append(letters, input.Text())
	}

	for numUserErr < maxUserErrors {
		draw.ClearTerminal()

		err := draw.Rendering(numUserErr)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Println()

		fmt.Printf("Слово: %s\n", string(secretWord))
		fmt.Printf("Ошибки(%d): %s \n", numUserErr, wrongLetters)
		fmt.Print("Буква: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		char = scanner.Text()

		err = CheckUserChar(letters, &char)

		if err != nil {
			fmt.Println("Введен некоректный символ. Попробуйте снова.")
			time.Sleep(numSecondsWaiting * time.Second)
			continue
		}

		var index int

		flagErr = false
		for _, runeValue := range word {
			for _, runeChar := range char {
				if runeChar == runeValue {
					secretWord[index] = runeChar
					flagErr = true
					break
				}
			}
			index += stepWord
		}

		for i := 0; i < len(wrongLetters); i++ {
			if wrongLetters[i] == char {
				flagErr = true
			}
		}

		if !flagErr {
			wrongLetters = append(wrongLetters, char)
			numUserErr += 1
		}

		if reflect.DeepEqual([]rune(expectedWord), secretWord) {
			draw.ClearTerminal()
			err = draw.Rendering(numUserErr)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			fmt.Printf("Слово: %s\n", string(secretWord))
			fmt.Println("Поздравляем, вы отгадали слово!")
			os.Exit(0)
		}
	}

	draw.ClearTerminal()
	err = draw.Rendering(numUserErr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("Вы проиграли :(")
}

func CheckUserChar(letters []string, inputChar *string) error {
	for _, letter := range letters {
		if strings.EqualFold(letter, *inputChar) {
			*inputChar = strings.ToLower(*inputChar)
			return nil
		}
	}

	return fmt.Errorf("invalid symbol")
}
