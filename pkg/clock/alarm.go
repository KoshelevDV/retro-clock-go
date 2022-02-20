package clock

import "github.com/koshelevdv/retro-clock-go/internal/placeholder"

func alarm(delay, duration int) {
	alarm := [...]placeholder.Placeholder{
		placeholder.Colon[2],
		placeholder.Characters[0],
		placeholder.Characters[1],
		placeholder.Characters[0],
		placeholder.Characters[2],
		placeholder.Characters[3],
		placeholder.Characters[4],
		placeholder.Colon[2],
	}

	if delay%10 >= 0 && delay%10 < duration {
		placeholder.Print(alarm[:], placeholder.Colon[:], delay, false)
	}
}
