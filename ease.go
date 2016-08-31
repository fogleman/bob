package bob

import "math"

func easeInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		return (t-1)*(2*t-2)*(2*t-2) + 1
	}
}

func easeOutElastic(t float64) float64 {
	p := 0.2
	return math.Pow(2, -10*t)*math.Sin((t-p/4)*(2*math.Pi)/p) + 1
}

func easeInBack(t float64) float64 {
	s := 1.70158
	return 1 - t*t*((s+1)*t-s)
}

func easeInQuint(t float64) float64 {
	return t * t * t * t * t
}
