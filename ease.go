package bob

import "math"

type EaseFunction func(float64) float64

func EaseLinear(t float64) float64 {
	return t
}

func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	} else {
		t = t*2 - 1
		return -0.5 * (t*(t-2) - 1)
	}
}

func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		return (t-1)*(2*t-2)*(2*t-2) + 1
	}
}

func EaseOutElastic(t float64) float64 {
	p := 0.2
	return math.Pow(2, -10*t)*math.Sin((t-p/4)*(2*math.Pi)/p) + 1
}

func EaseInBack(t float64) float64 {
	s := 1.70158
	return 1 - t*t*((s+1)*t-s)
}

func EaseInQuint(t float64) float64 {
	return t * t * t * t * t
}
