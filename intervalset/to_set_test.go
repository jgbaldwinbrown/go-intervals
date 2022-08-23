package intervalset

import (
	"testing"
	"fmt"
)

func TestToSet(t *testing.T) {
	spans := []Span{
		Span{5,8},
		Span{7,10},
		Span{22, 33},
	}
	spanptrs := make([]*Span, len(spans))
	for i, _ := range spans {
		spanptrs[i] = &spans[i]
	}
	set := ImmSet(spanptrs)
	fmt.Println(set)
}

func TestToSetVal(t *testing.T) {
	vspans := []Vspan{
		Vspan{5,8},
		Vspan{7,10},
		Vspan{22, 33},
	}
	set := ImmSet(vspans)
	fmt.Println(set)
}
