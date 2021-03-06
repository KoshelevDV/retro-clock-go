package placeholder

import "fmt"

func Print(content []Placeholder, colon []Placeholder, sec int, blink bool) {
	for line := range content[0] {
		for index, digit := range content {
			next := content[index][line]
			if blink {
				if Contains(colon, digit) && sec%2 == 0 {
					next = "   "
				}
			}
			fmt.Print(next, "  ")
		}
		fmt.Println()
	}
}

func CyclePrint(content []Placeholder, colon []Placeholder, sec int, blink bool) {
	for i := 0; i < len(content); i++ {
		Print(content[i:], colon, sec, blink)
	}
}
