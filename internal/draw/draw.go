package draw

import (
	"fmt"
)

var picture = []string{
	" ___\n    |\n    |\n    |\n ___|_",
	" ___\n o  |\n    |\n    |\n ___|_",
	" ___\n o  |\n |  |\n    |\n ___|_",
	" ___\n o  |\n/|  |\n    |\n ___|_",
	" ___\n o  |\n/|\\ |\n    |\n ___|_",
	" ___\n o  |\n/|\\ |\n/   |\n ___|_",
	" ___\n o  |\n/|\\ |\n/ \\ |\n ___|_",
}

func Rendering(numPicture int) error {

	if numPicture > len(picture)-1 {
		return fmt.Errorf("slice out of bounds")
	}

	fmt.Println(picture[numPicture])
	return nil
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
