package intervalset

import (
	"fmt"
)

type Span struct {
	Min, Max int
}

func MakeZeroSpan() Interval {
	return Zero()
}

// Cast returns a *Span from an Interval interface, or it panics.
func Cast(i Interval) *Span {
	x, ok := i.(*Span)
	if !ok {
		panic(fmt.Errorf("interval must be an Span: %v", i))
	}
	return x
}

// Zero returns the Zero value for Span.
func Zero() *Span {
	return &Span{}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Span) MakeZero() Interval {
	return MakeZeroSpan()
}

func (s *Span) String() string {
	return fmt.Sprintf("[%d, %d)", s.Min, s.Max)
}

func (s *Span) Equal(t *Span) bool {
	return s.Min == t.Min && s.Max == t.Max
}

// Intersect returns the intersection of an interval with another
// interval. The function may panic if the other interval is incompatible.
func (s *Span) Intersect(tInt Interval) Interval {
	t := Cast(tInt)
	result := &Span{
		Max(s.Min, t.Min),
		Min(s.Max, t.Max),
	}
	if result.Min < result.Max {
		return result
	}
	return Zero()
}

// Before returns true if the interval is completely before another interval.
func (s *Span) Before(tInt Interval) bool {
	t := Cast(tInt)
	return s.Max <= t.Min
}

// IsZero returns true for the Zero value of an interval.
func (s *Span) IsZero() bool {
	return s.Min == 0 && s.Max == 0
}

// Bisect returns two intervals, one on either lower side of x and one on the
// upper side of x, corresponding to the subtraction of x from the original
// interval. The returned intervals are always within the range of the
// original interval.
func (s *Span) Bisect(tInt Interval) (Interval, Interval) {
	intersection := Cast(s.Intersect(tInt))
	if intersection.IsZero() {
		if s.Before(tInt) {
			return s, Zero()
		}
		return Zero(), s
	}
	maybeZero := func(Min, Max int) *Span {
		if Min == Max {
			return Zero()
		}
		return &Span{Min, Max}
	}
	return maybeZero(s.Min, intersection.Min), maybeZero(intersection.Max, s.Max)

}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the Zero interval if they are not.
func (s *Span) Adjoin(tInt Interval) Interval {
	t := Cast(tInt)
	if s.Max == t.Min {
		return &Span{s.Min, t.Max}
	}
	if t.Max == s.Min {
		return &Span{t.Min, s.Max}
	}
	return Zero()
}

// Encompass returns an interval that covers the exact extents of two
// intervals.
func (s *Span) Encompass(tInt Interval) Interval {
	t := Cast(tInt)
	return &Span{Min(s.Min, t.Min), Max(s.Max, t.Max)}
}

func AllIntervals(s SetInput) []*Span {
	result := []*Span{}
	s.IntervalsBetween(s.Extent(), func(x Interval) bool {
		result = append(result, Cast(x))
		return true
	})
	return result
}

