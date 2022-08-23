package intervalset

import (
	"testing"
	"fmt"
)

func TestSettable(t *testing.T) {
	vspans1 := []Vspan{
		Vspan{5,8},
		Vspan{7,10},
		Vspan{22, 33},
	}
	vspans2 := []Vspan{
		Vspan{4,6},
		Vspan{20,25},
	}
	out := Intersect(vspans1, vspans2)
	fmt.Println(out)
}
