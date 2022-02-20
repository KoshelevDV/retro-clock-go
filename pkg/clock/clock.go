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

func CycleClock() {
	i := 0
	toLeft := true

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

		temp := [...]placeholder.Placeholder{
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
			placeholder.Colon[2],
		}

		if toLeft {
			placeholder.Print(append(clock[:], temp[:]...)[i:i+10], placeholder.Colon[:], sec, true)
		} else {
			placeholder.Print(append(temp[:], clock[:]...)[i:i+10], placeholder.Colon[:], sec, true)
		}
		func() {
			fmt.Println()
			alarm(sec, 2)
		}()

		if (milsec/100000000)%10 == 0 {
			i += 1
		}
		if i >= 10 && toLeft {
			i = 0
			toLeft = !toLeft
		} else if i >= 9 && !toLeft {
			i = 0
			toLeft = !toLeft
		}
		time.Sleep(time.Millisecond * 100)
	}
}
