package intervalset

type Settable interface {
	ImmutableSet() *ImmutableSet
	FromSetInput(SetInput)
}

func Contains[T Settable](vals T, ival Interval) bool {
	s := vals.ImmutableSet()
	return s.Contains(ival)
}

func Extent[T Settable](vals T) Interval {
	s := vals.ImmutableSet()
	return s.Extent()
}

func Intersect[T Settable](v1, v2 T) T {
	s1, s2 := v1.ImmutableSet(), v2.ImmutableSet()
	sout := s1.Intersect(s2)

	var vout T
	vout.FromSetInput(sout)
	return vout
}

func Intervals[T Settable](vals T, f IntervalReceiver) {
	s := vals.ImmutableSet()
	s.Intervals(f)
}

func IntervalsBetween[T Settable](vals T, extents Interval, f IntervalReceiver) {
	s := vals.ImmutableSet()
	s.IntervalsBetween(extents, f)
}

func Sub[T Settable](v1, v2 T) T {
	s1, s2 := v1.ImmutableSet(), v2.ImmutableSet()
	sout := s1.Sub(s2)

	var vout T
	vout.FromSetInput(sout)
	return vout
}

func Union[T Settable](v1, v2 T) T {
	s1, s2 := v1.ImmutableSet(), v2.ImmutableSet()
	sout := s1.Union(s2)

	var vout T
	vout.FromSetInput(sout)
	return vout
}

/*
Jim@T480:intervalset$ go doc ImmutableSet
package intervalset // import "."

type ImmutableSet struct {
	// Has unexported fields.
}
    ImmutableSet is a set of interval objects. It provides various set theory
    operations.

func NewImmutableSet(intervals []Interval) *ImmutableSet
func NewImmutableSetV1(intervals []Interval, makeZero func() Interval) *ImmutableSet
func (s *ImmutableSet) Contains(ival Interval) bool
func (s *ImmutableSet) Extent() Interval
func (s *ImmutableSet) Intersect(b SetInput) *ImmutableSet
func (s *ImmutableSet) Intervals(f IntervalReceiver)
func (s *ImmutableSet) IntervalsBetween(extents Interval, f IntervalReceiver)
func (s *ImmutableSet) String() string
func (s *ImmutableSet) Sub(b SetInput) *ImmutableSet
func (s *ImmutableSet) Union(b SetInput) *ImmutableSet
*/
