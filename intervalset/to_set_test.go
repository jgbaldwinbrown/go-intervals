package intervalset

import (
	"testing"
	"fmt"
)

// func TestValToSet(t *testing.T) {
// 	spans := []Span{
// 		Span{5,8},
// 		Span{7,10},
// 		Span{22, 33},
// 	}
// 	fmt.Println(ToNewSet(spans))
// }

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
	set := ToSetV1(spanptrs, MakeZeroSpan)
	fmt.Println(set)
}

func TestToSetPtrs(t *testing.T) {
	spans := []Span{
		Span{5,8},
		Span{7,10},
		Span{22, 33},
	}
	set := ToPtrSet(spans)
	fmt.Println(set)
}

func TestToSetVal(t *testing.T) {
	vspans := []Vspan{
		Vspan{5,8},
		Vspan{7,10},
		Vspan{22, 33},
	}
	set := ToSet(vspans)
	fmt.Println(set)
}
