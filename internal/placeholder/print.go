package placeholder

import "fmt"

func Print(content []Placeholder) {
	for line := range content[0] {
		for digit := range content {
			fmt.Print(content[digit][line], "  ")
		}
		fmt.Println()
	}
}
