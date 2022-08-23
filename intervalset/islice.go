package intervalset

type Zinterval interface {
	Interval
	MakeZero() Interval
}

type setInput struct {
	I Interval
}

func (s setInput) Extent() Interval { return s.I }

func (s setInput) IntervalsBetween(extent Interval, f IntervalReceiver) {
	f(s.I.Intersect(extent))
}

func ImmSet[T Zinterval](vals []T) *ImmutableSet {
	zf := func() Interval {
		var t T
		return t.MakeZero()
	}
	set := NewSetV1([]Interval{}, zf)
	for _, val := range vals {
		set.Add(setInput{Interval(val)})
	}
	return set.ImmutableSet()
}

func ToZslice[T Zinterval](set *ImmutableSet) []T {
	var out []T
	f := func(i Interval) bool {
		out = append(out, i.(T))
		return true
	}
	set.Intervals(f)
	return out
}
