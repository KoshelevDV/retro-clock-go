package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/koshelevdv/retro-clock-go/internal/digits"
)

func main() {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour := now.Hour()
		min := now.Minute()
		sec := now.Second()

		clock := [...]digits.Placeholder{
			digits.Digits[hour/10], digits.Digits[hour%10],
			digits.Colon[sec%2],
			digits.Digits[min/10], digits.Digits[min%10],
			digits.Colon[sec%2],
			digits.Digits[sec/10], digits.Digits[sec%10],
		}
		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
