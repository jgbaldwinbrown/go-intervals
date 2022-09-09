package intervalset

type setInput struct {
	I Interval
}

func (s setInput) Extent() Interval { return s.I }

func (s setInput) IntervalsBetween(extent Interval, f IntervalReceiver) {
	f(s.I.Intersect(extent))
}

func ToSet[T Interval](vals []T) *Set {
	zf := func() Interval {
		var t T
		return t
	}
	set := NewSetV1([]Interval{}, zf)
	for _, val := range vals {
		set.Add(setInput{Interval(val)})
	}
	return set
}

func ImmSet[T Interval](vals []T) *ImmutableSet {
	set := ToSet(vals)
	return set.ImmutableSet()
}

func Slice[T Interval](set *ImmutableSet) []T {
	var out []T
	f := func(i Interval) bool {
		out = append(out, i.(T))
		return true
	}
	set.Intervals(f)
	return out
}
