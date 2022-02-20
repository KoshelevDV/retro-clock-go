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
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder.Placeholder{
			placeholder.Digits[hour/10], placeholder.Digits[hour%10],
			placeholder.Colon[sec%2],
			placeholder.Digits[min/10], placeholder.Digits[min%10],
			placeholder.Colon[sec%2],
			placeholder.Digits[sec/10], placeholder.Digits[sec%10],
		}

		placeholder.Print(clock[:])
		func() {
			fmt.Println()
			alarm(sec, 2)
		}()

		time.Sleep(time.Second)
	}
}

func alarm(delay, duration int) {
	alarm := [...]placeholder.Placeholder{
		placeholder.Colon[1],
		placeholder.Characters[0],
		placeholder.Characters[1],
		placeholder.Characters[0],
		placeholder.Characters[2],
		placeholder.Characters[3],
		placeholder.Characters[4],
		placeholder.Colon[1],
	}
	if delay%10 >= 0 && delay%10 <= duration {
		placeholder.Print(alarm[:])
	}

}
