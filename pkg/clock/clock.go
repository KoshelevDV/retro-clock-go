package clock

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/koshelevdv/retro-clock-go/internal/placeholder"
)

func Clock() {
	for shift := 0; ; shift++ {
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

		for line := range clock[0] {
			l := len(clock)
			s, e := shift%l, l

			if shift%(l*2) >= l {
				s, e = 0, s
			}

			for j := 0; j < l-e; j++ {
				fmt.Print("     ")
			}

			for i := s; i < e; i++ {
				next := clock[i][line]
				if placeholder.Contains(placeholder.Colon[:], clock[i]) && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Millisecond * 1000)
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
