package draw

import "fmt"

func Rendering(e int) {
	if e == 0 {
		fmt.Println(" ___")
		fmt.Println("    |")
		fmt.Println("    |")
		fmt.Println("    |")
		fmt.Println(" ___|_")
	} else if e == 1 {
		fmt.Println(" ___")
		fmt.Println(" o  |")
		fmt.Println("    |")
		fmt.Println("    |")
		fmt.Println(" ___|_")
	} else if e == 2 {
		fmt.Println(" ___")
		fmt.Println(" o  |")
		fmt.Println(" |  |")
		fmt.Println("    |")
		fmt.Println(" ___|_")
	} else if e == 3 {
		fmt.Println(" ___")
		fmt.Println(" o  |")
		fmt.Println("/|  |")
		fmt.Println("    |")
		fmt.Println(" ___|_")
	} else if e == 4 {
		fmt.Println(" ___")
		fmt.Println(" o  |")
		fmt.Println("/|\\ |")
		fmt.Println("    |")
		fmt.Println(" ___|_")
	} else if e == 5 {
		fmt.Println(" ___")
		fmt.Println(" o  |")
		fmt.Println("/|\\ |")
		fmt.Println("/   |")
		fmt.Println(" ___|_")
	} else if e == 6 {
		fmt.Println(" ___")
		fmt.Println(" o  |")
		fmt.Println("/|\\ |")
		fmt.Println("/ \\ |")
		fmt.Println(" ___|_")
	}
}
