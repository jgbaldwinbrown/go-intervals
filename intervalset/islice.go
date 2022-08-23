package intervalset

type Zinterval interface {
	Interval
	MakeZero() Interval
}

type InterSlice[T Zinterval] struct {
	Slice []T
}

func (s *InterSlice[T]) ImmutableSet() *ImmutableSet {
	zf := func() Interval {
		var t T
		return t.MakeZero()
	}
	return ToSetV1(s.Slice, zf).ImmutableSet()
}

func (s *InterSlice[T]) FromImmSet(set *ImmutableSet) {
	f := func(i Interval) bool {
		s.Slice = append(s.Slice, i.(T))
		return true
	}
	set.Intervals(f)
}
