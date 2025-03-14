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

<<<<<<< HEAD
	game.GameLogic(lettersFile, secretWord, expectedWord, word)
=======
	filename := "list_of_letters.txt"

	for usrErr < 6 {
		file, e := os.Open(filename)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при открытии файла!\n%v", e)
		}

		ClearTerminal()

		draw.Rendering(usrErr)
		fmt.Println()

		fmt.Printf("Слово: %s\n", string(lenWord))
		fmt.Printf("Ошибки(%d): %s \n", usrErr, s)
		fmt.Print("Буква: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		char = scanner.Text()

		err := CheckUserChar(file, &char)

		if err != nil {
			fmt.Println("Введен некоректный символ. Попробуйте снова.")
			time.Sleep(2 * time.Second)
			continue
		}

		var index int

		flagErr = false
		for _, runeValue := range word {
			for _, runeChar := range char {
				if runeChar == runeValue {
					lenWord[index] = runeChar
					flagErr = true
					break
				}
			}
			index += 2
		}

		for i := 0; i < len(s); i++ {
			if s[i] == char {
				flagErr = true
			}
		}

		if !flagErr {
			s = append(s, char)
			usrErr += 1
		}

		if string(lenWord) == test_word {
			ClearTerminal()
			draw.Rendering(usrErr)
			fmt.Printf("Слово: %s\n", string(lenWord))
			fmt.Println("Поздравляем, вы отгадали слово!")
			os.Exit(0)
		}
	}

	ClearTerminal()
	draw.Rendering(usrErr)
	fmt.Println("Вы проиграли :(")
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func CheckUserChar(f *os.File, c *string) error {
	input := bufio.NewScanner(f)
	a := false

	for input.Scan() {
		if strings.EqualFold(input.Text(), *c) {
			*c = strings.ToLower(*c)
			a = true
			break
		}
	}

	f.Close()

	if !a {
		return fmt.Errorf("введен некорректный символ")
	}

	return nil
>>>>>>> 0a0eae41b8898de5a9258ffa556057c490a47db8
}
