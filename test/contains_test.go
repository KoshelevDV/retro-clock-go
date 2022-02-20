package placeholder

import (
	"testing"

	"github.com/koshelevdv/retro-clock-go/internal/placeholder"
)

func TestContains(t *testing.T) {
	nums := placeholder.Digits
	ttt := [5]string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	if !placeholder.Contains(nums[:], ttt) {
		t.Fatal()
	}
}
