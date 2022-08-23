package intervalset

import (
	"fmt"
)

type Vspan struct {
	Min, Max int
}

// Cast returns a Vspan from an Interval interface, or it panics.
func Vcast(i Interval) Vspan {
	x, ok := i.(Vspan)
	if !ok {
		panic(fmt.Errorf("interval must be an Vspan: %v", i))
	}
	return x
}

// Zero returns the Zero value for Vspan.
func Vzero() Vspan {
	return Vspan{}
}

func (s Vspan) String() string {
	return fmt.Sprintf("[%d, %d)", s.Min, s.Max)
}

func (s Vspan) Equal(t Vspan) bool {
	return s.Min == t.Min && s.Max == t.Max
}

// Intersect returns the intersection of an interval with another
// interval. The function may panic if the other interval is incompatible.
func (s Vspan) Intersect(tInt Interval) Interval {
	t := Vcast(tInt)
	result := Vspan{
		Max(s.Min, t.Min),
		Min(s.Max, t.Max),
	}
	if result.Min < result.Max {
		return result
	}
	return Vzero()
}

// Before returns true if the interval is completely before another interval.
func (s Vspan) Before(tInt Interval) bool {
	t := Vcast(tInt)
	return s.Max <= t.Min
}

// IsZero returns true for the Zero value of an interval.
func (s Vspan) IsZero() bool {
	return s.Min == 0 && s.Max == 0
}

// Bisect returns two intervals, one on either lower side of x and one on the
// upper side of x, corresponding to the subtraction of x from the original
// interval. The returned intervals are always within the range of the
// original interval.
func (s Vspan) Bisect(tInt Interval) (Interval, Interval) {
	intersection := Vcast(s.Intersect(tInt))
	if intersection.IsZero() {
		if s.Before(tInt) {
			return s, Vzero()
		}
		return Vzero(), s
	}
	maybeZero := func(Min, Max int) Vspan {
		if Min == Max {
			return Vzero()
		}
		return Vspan{Min, Max}
	}
	return maybeZero(s.Min, intersection.Min), maybeZero(intersection.Max, s.Max)

}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the Zero interval if they are not.
func (s Vspan) Adjoin(tInt Interval) Interval {
	t := Vcast(tInt)
	if s.Max == t.Min {
		return Vspan{s.Min, t.Max}
	}
	if t.Max == s.Min {
		return Vspan{t.Min, s.Max}
	}
	return Vzero()
}

// Encompass returns an interval that covers the exact extents of two
// intervals.
func (s Vspan) Encompass(tInt Interval) Interval {
	t := Vcast(tInt)
	return Vspan{Min(s.Min, t.Min), Max(s.Max, t.Max)}
}

func AllVintervals(s SetInput) []Vspan {
	result := []Vspan{}
	s.IntervalsBetween(s.Extent(), func(x Interval) bool {
		result = append(result, Vcast(x))
		return true
	})
	return result
}

