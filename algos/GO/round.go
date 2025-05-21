package algos

func round(x float64, n int) float64 {
	// round x to n decimal places
	if n < 0 {
		return x
	}
	m := 1.0
	for i := 0; i < n; i++ {
		m *= 10.0
	}
	return float64(int(x*m+0.5)) / m
}
