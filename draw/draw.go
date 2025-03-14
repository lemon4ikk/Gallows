package draw

import "fmt"

func Rendering(numPicture int) error {
	picture := []string{
		" ___\n    |\n    |\n    |\n ___|_",
		" ___\n o  |\n    |\n    |\n ___|_",
		" ___\n o  |\n |  |\n    |\n ___|_",
		" ___\n o  |\n/|  |\n    |\n ___|_",
		" ___\n o  |\n/|\\ |\n    |\n ___|_",
		" ___\n o  |\n/|\\ |\n/   |\n ___|_",
		" ___\n o  |\n/|\\ |\n/ \\ |\n ___|_",
	}

	if numPicture > len(picture)-1 {
		return fmt.Errorf("slice out of bounds")
	}

	fmt.Println(picture[numPicture])
	return nil
}
