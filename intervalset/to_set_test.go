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
	set := ToNewSet(spanptrs, MakeZeroSpan)
	fmt.Println(set)
}

func TestToSetPtrs(t *testing.T) {
	spans := []Span{
		Span{5,8},
		Span{7,10},
		Span{22, 33},
	}
	spanptrs := ToPtrInterval(spans)
	set := ToNewSet(spanptrs, MakeZeroSpan)
	fmt.Println(set)
}
