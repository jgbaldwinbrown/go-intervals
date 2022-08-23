package intervalset

import (
	"fmt"
)

type Bspan struct {
	Lchr, Rchr string
	Min, Max int
}

func Smin(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func Smax(a, b string) string {
	if a > b {
		return a
	}
	return b
}


// Cast returns a Bspan from an Interval interface, or it panics.
func Bcast(i Interval) Bspan {
	x, ok := i.(Bspan)
	if !ok {
		panic(fmt.Errorf("interval must be an Bspan: %v", i))
	}
	return x
}

// Zero returns the Zero value for Bspan.
func Bzero() Bspan {
	return Bspan{}
}

func (s Bspan) Chreq(t Bspan) bool {
	return s.Lchr == t.Lchr && s.Rchr == t.Rchr
}

func (s Bspan) String() string {
	return fmt.Sprintf("[%d, %d)", s.Min, s.Max)
}

func (s Bspan) Equal(t Bspan) bool {
	return s.Chreq(t) && s.Min == t.Min && s.Max == t.Max
}

// Intersect returns the intersection of an interval with another
// interval. The function may panic if the other interval is incompatible.
func (s Bspan) Intersect(tInt Interval) Interval {
	t := Bcast(tInt)
	result := Bspan{
		Smin(s.Lchr, t.Lchr),
		Smax(s.Rchr, t.Rchr),
		Max(s.Min, t.Min),
		Min(s.Max, t.Max),
	}
	if result.Min < result.Max && s.Chreq(t) {
		return result
	}
	return Bzero()
}

// Before returns true if the interval is completely before another interval.
func (s Bspan) Before(tInt Interval) bool {
	t := Bcast(tInt)
	return s.Max <= t.Min && s.Rchr <= t.Lchr
}

// IsZero returns true for the Zero value of an interval.
func (s Bspan) IsZero() bool {
	return s.Min == 0 && s.Max == 0 && s.Lchr == "" && s.Rchr == ""
}

// Bisect returns two intervals, one on either lower side of x and one on the
// upper side of x, corresponding to the subtraction of x from the original
// interval. The returned intervals are always within the range of the
// original interval.
func (s Bspan) Bisect(tInt Interval) (Interval, Interval) {
	intersection := Bcast(s.Intersect(tInt))
	if intersection.IsZero() {
		if s.Before(tInt) {
			return s, Bzero()
		}
		return Bzero(), s
	}
	maybeZero := func(Cmin, Cmax string, Min, Max int) Bspan {
		if Min == Max && Cmin == Cmax {
			return Bzero()
		}
		return Bspan{Cmin, Cmax, Min, Max}
	}
	return maybeZero(s.Lchr, intersection.Lchr, s.Min, intersection.Min), maybeZero(s.Rchr, intersection.Rchr, intersection.Max, s.Max)

}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the Zero interval if they are not.
func (s Bspan) Adjoin(tInt Interval) Interval {
	t := Bcast(tInt)
	if s.Max == t.Min && s.Rchr == t.Lchr {
		return Bspan{s.Lchr, t.Rchr, s.Min, t.Max}
	}
	if t.Max == s.Min && s.Lchr == t.Rchr {
		return Bspan{t.Lchr, s.Rchr, t.Min, s.Max}
	}
	return Bzero()
}

// Encompass returns an interval that covers the exact extents of two
// intervals.
func (s Bspan) Encompass(tInt Interval) Interval {
	t := Bcast(tInt)
	return Bspan{Smin(s.Lchr, t.Lchr), Smax(s.Rchr, t.Rchr), Min(s.Min, t.Min), Max(s.Max, t.Max)}
}

func AllBintervals(s SetInput) []Bspan {
	result := []Bspan{}
	s.IntervalsBetween(s.Extent(), func(x Interval) bool {
		result = append(result, Bcast(x))
		return true
	})
	return result
}

