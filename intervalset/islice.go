package intervalset

type Zinterval interface {
	Interval
	MakeZero() Interval
}

func ImmSet[T Zinterval](vals []T) *ImmutableSet {
	zf := func() Interval {
		var t T
		return t.MakeZero()
	}
	return ToSetV1(vals, zf).ImmutableSet()
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
