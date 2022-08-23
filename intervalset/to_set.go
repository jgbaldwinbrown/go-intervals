package intervalset

type setInput struct {
	I Interval
}

func (s setInput) Extent() Interval { return s.I }

func (s setInput) IntervalsBetween(extent Interval, f IntervalReceiver) {
	f(s.I.Intersect(extent))
}

func ToInterval[T Interval](concretes []T) []Interval {
	inters := make([]Interval, len(concretes))
	for i, c := range concretes {
		inters[i] = Interval(c)
	}
	return inters
}

func ToPtrInterval[T any](concretes []T) []Interval {
	inters := make([]Interval, len(concretes))
	for i, c := range concretes {
		p := new(T)
		*p = c
		inters[i] = any(p).(Interval)
	}
	return inters
}

func ToNewSet[T Interval](concretes []T, f func() Interval) *Set {
	set := NewSetV1([]Interval{}, f)
	inters := ToInterval(concretes)
	for _, inter := range inters {
		set.Add(setInput{inter})
	}
	return set
}
