package funcs

import (
	"math"
	"sort"
)

func LinReg(pop []float64) (float64, float64) {
	var a, b float64
	c := float64(len(pop) / 2)
	s := make([]float64, 0)
	for i, v := range pop {
		s = append(s, float64(i))
		a += (float64(i) - c) * (v - Average(pop))
	}
	a = a / float64(len(pop))
	a = a / Variance(s)
	b = Average(pop) - a*Average(s)
	return a, b
}

func CorCoeff(pop []float64) float64 {
	var a float64
	c := float64(len(pop) / 2)
	s := make([]float64, 0)
	for i, v := range pop {
		s = append(s, float64(i))
		a += (float64(i) - c) * (v - Average(pop))
	}
	a = a / float64(len(pop))
	res := a / ((Deviation(pop)) * Deviation(s))
	return res
}

func Average(pop []float64) float64 {
	var sum float64
	for _, v := range pop {
		sum += v
	}
	return sum / float64(len(pop))
}
func Median(pop []float64) float64 {
	l := len(pop) / 2
	sort.Float64s(pop)
	if len(pop)%2 != 0 {
		return pop[l]
	}
	return (pop[l] + pop[l-1]) / 2
}

func Variance(pop []float64) float64 {
	var s float64
	for _, v := range pop {
		s += (v - Average(pop)) * (v - Average(pop))
	}
	return (s) / float64(len(pop))
}

func Deviation(pop []float64) float64 {
	return math.Sqrt(Variance(pop))
}
