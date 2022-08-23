package intervalset

import (
	"testing"
	"fmt"
)

func TestBspan(t *testing.T) {
	bspans1 := []Bspan{
		Bspan{"1", "1", 5,8},
		Bspan{"1", "1", 7,10},
		Bspan{"2", "2", 22, 33},
	}
	bspans2 := []Bspan{
		Bspan{"1", "1", 4,6},
		Bspan{"2", "2", 20,25},
	}
	out := Intersect(bspans1, bspans2)
	fmt.Println(out)
}
