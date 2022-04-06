package client

import "ch15/series"
import "testing"

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.square(5))
}
