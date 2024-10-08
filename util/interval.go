package util

type Interval struct {
	min float64
	max float64
}

func NewInterval(min, max float64) *Interval {
	return &Interval{min: min, max: max}
}

func (interval *Interval) Min() float64 {
	return interval.min
}

func (interval *Interval) SetMin(min float64) {
	interval.min = min
}

func (interval *Interval) Max() float64 {
	return interval.max
}

func (interval *Interval) SetMax(max float64) {
	interval.max = max
}

func (interval *Interval) Contains(x float64) bool {
	return interval.min <= x && x <= interval.max
}

func (interval *Interval) Surrounds(x float64) bool {
	return interval.min < x && x < interval.max
}

func (interval *Interval) Clamp(x float64) float64 {
	switch {
	case x < interval.min:
		return interval.min
	case x > interval.max:
		return interval.max
	default:
		return x
	}
}
