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

func ToSetV1[T Interval](concretes []T, f func() Interval) *Set {
	set := NewSetV1([]Interval{}, f)
	inters := ToInterval(concretes)
	for _, inter := range inters {
		set.Add(setInput{inter})
	}
	return set
}

func ToPtrSetV1[T any](concretes []T, f func() Interval) *Set {
	set := NewSetV1([]Interval{}, f)
	inters := ToPtrInterval(concretes)
	for _, inter := range inters {
		set.Add(setInput{inter})
	}
	return set
}

func ToSet[T Interval](concretes []T) *Set {
	f := func() Interval { var t T; return t }
	return ToSetV1(concretes, f)
}

func ToPtrSet[T any](concretes []T) *Set {
	f := func() Interval {
		t := new(T)
		return any(t).(Interval)
	}
	return ToPtrSetV1(concretes, f)
}
