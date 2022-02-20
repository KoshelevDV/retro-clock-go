package clock

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/koshelevdv/retro-clock-go/internal/placeholder"
)

func Clock() {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec, milsec := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()

		clock := [...]placeholder.Placeholder{
			placeholder.Digits[hour/10], placeholder.Digits[hour%10],
			placeholder.Colon[0],
			placeholder.Digits[min/10], placeholder.Digits[min%10],
			placeholder.Colon[0],
			placeholder.Digits[sec/10], placeholder.Digits[sec%10],
			placeholder.Colon[1],
			placeholder.Digits[milsec/100000000],
		}

		placeholder.Print(clock[:], placeholder.Colon[:], sec, true)
		func() {
			fmt.Println()
			alarm(sec, 2)
		}()

		time.Sleep(time.Millisecond * 100)
	}
}
