package game

import (
	"bufio"
	"fmt"
	"os"
)

const (
	startCommand = "Начать новую игру"
	exitCommand  = "Выйти"
)

func Menu() {
	var userMsg string

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
