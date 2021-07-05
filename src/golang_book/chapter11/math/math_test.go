package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var testsAverage = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

var testsMin = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{1, 2, 3, 4, 5, 6}, 1},
	{[]float64{-1, 1}, -1},
	{[]float64{}, 0},
}

var testsMax = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{1, 2, 3, 4, 5, 6}, 6},
	{[]float64{-1, 1}, 1},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range testsAverage {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range testsMin {
		v := Min(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range testsMax {
		v := Max(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
