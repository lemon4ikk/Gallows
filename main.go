package main

import (
	"bufio"
	"fmt"
	"gallows/draw"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	var userMsg string
	start_message := "Начать новую игру"
	out_message := "Выйти"

	fmt.Println("Для начала игры введите: " + start_message)
	fmt.Println("Для выхода из игры введите: " + out_message)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userMsg = scanner.Text()

	if userMsg == out_message {
		os.Exit(0)
	}

	if userMsg == start_message {
		Start()
	}
}

func Start() {
	dictionaryFile := "dictionary.txt"
	var dictionary []string
	file, e := os.Open(dictionaryFile)

	if e != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при открытии файла!\n%v", e)
	}

	scanWord := bufio.NewScanner(file)

	for scanWord.Scan() {
		dictionary = append(dictionary, scanWord.Text())
	}

	index := rand.Intn(len(dictionary))
	file.Close()

	word := dictionary[index]
	var test_word string

	for i := 0; i < utf8.RuneCountInString(word); i++ {
		test_word = test_word + "_" + " "
	}

	lenWord := make([]rune, utf8.RuneCountInString(test_word))
	var usrErr int
	var flagErr bool
	var char string
	s := []string{}

	for i, k := range test_word {
		lenWord[i] = k
	}

	test_word = ""

	for _, char := range word {
		test_word += string(char) + " "
	}

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
}
