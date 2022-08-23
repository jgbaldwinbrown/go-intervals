package intervalset

func Contains[T Interval](vals []T, ival Interval) bool {
	s := ImmSet(vals)
	return s.Contains(ival)
}

func Extent[T Interval](vals []T) Interval {
	s := ImmSet(vals)
	return s.Extent()
}

func Intersect[T Interval](v1, v2 []T) []T {
	s1 := ImmSet(v1)
	s2 := ImmSet(v2)
	sout := s1.Intersect(s2)

	return ToZslice[T](sout)
}

func Intervals[T Interval](vals []T, f IntervalReceiver) {
	s := ImmSet(vals)
	s.Intervals(f)
}

func IntervalsBetween[T Interval](vals []T, extents Interval, f IntervalReceiver) {
	s := ImmSet(vals)
	s.IntervalsBetween(extents, f)
}

func Sub[T Interval](v1, v2 []T) []T {
	s1 := ImmSet(v1)
	s2 := ImmSet(v2)
	sout := s1.Sub(s2)

	return ToZslice[T](sout)
}

func Union[T Interval](v1, v2 []T) []T {
	s1 := ImmSet(v1)
	s2 := ImmSet(v2)
	sout := s1.Union(s2)

	return ToZslice[T](sout)
}
