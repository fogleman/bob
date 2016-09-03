package bob

type Float64Item struct {
	T0, T1   float64
	X0, X1   float64
	Function EaseFunction
}

type Float64Variable []Float64Item

func (v *Float64Variable) Set(t, x float64) {
	*v = append(*v, Float64Item{t, t, x, x, nil})
}

func (v *Float64Variable) Add(t0, t1, x0, x1 float64, f EaseFunction) {
	*v = append(*v, Float64Item{t0, t1, x0, x1, f})
}

func (v *Float64Variable) Get(t float64) float64 {
	var x float64
	for _, item := range *v {
		if t >= item.T0 && t < item.T1 {
			u := (t - item.T0) / (item.T1 - item.T0)
			p := item.Function(u)
			return item.X0 + (item.X1-item.X0)*p
		}
		if t >= item.T1 {
			x = item.X1
		}
	}
	return x
}

func (v *Float64Variable) Changed(t0, t1 float64) bool {
	if t0 < 0 || t1 < 0 {
		return true
	}
	return v.Get(t0) != v.Get(t1)
}
