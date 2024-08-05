package main

import (
	"math"
	"testing"
)

func TestCheckDataQuality(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	var err error // Declare the err variable
	err = checkDataQuality(x, y)
	if err != nil {
		t.Errorf("checkDataQuality() returned an error for valid data: %v", err)
	}

	// Test case 2: Empty data
	x = []float64{}
	y = []float64{}
	err = checkDataQuality(x, y)
	if err == nil {
		t.Error("checkDataQuality() did not return an error for empty data")
	}
}

func TestXYData(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	xy := xyData(x, y)
	if len(xy) != len(x) {
		t.Errorf("xyData() returned an incorrect number of data points: expected %d, got %d", len(x), len(xy))
	}

	// Test case 2: Empty data
	x = []float64{}
	y = []float64{}
	xy = xyData(x, y)
	if len(xy) != 0 {
		t.Errorf("xyData() did not return an empty XYs for empty data: got %d data points", len(xy))
	}
}

func TestLinearRegression(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	series, err := linearRegression(x, y)
	if err != nil {
		t.Errorf("linearRegression() returned an error for valid data: %v", err)
	}
	if len(series) != 2 {
		t.Errorf("linearRegression() returned an incorrect number of series: expected 2, got %d", len(series))
	}

	// Test case 2: Empty data
	x = []float64{}
	y = []float64{}
	series, err = linearRegression(x, y)
	if err == nil {
		t.Error("linearRegression() did not return an error for empty data")
	}
	if len(series) != 0 {
		t.Errorf("linearRegression() did not return an empty series for empty data: got %d series", len(series))
	}
}

func TestCalculateRSquared(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	intercept := 0.0
	slope := 2.0
	rSquared, err := calculateRSquared(x, y, intercept, slope)
	if err != nil {
		t.Errorf("calculateRSquared() returned an error for valid data: %v", err)
	}
	if math.Abs(rSquared-1.0) > 1e-6 {
		t.Errorf("calculateRSquared() returned an incorrect R-squared value: expected 1.0, got %f", rSquared)
	}

	// Test case 2: Empty data
	x = []float64{}
	y = []float64{}
	intercept = 0.0
	slope = 0.0
	rSquared, err = calculateRSquared(x, y, intercept, slope)
	if err == nil {
		t.Error("calculateRSquared() did not return an error for empty data")
	}
	if math.Abs(rSquared-0.0) > 1e-6 {
		t.Errorf("calculateRSquared() did not return 0.0 for empty data: got %f", rSquared)
	}
}

func TestCreateScatterPlot(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	var err error
	createScatterPlot(1, x, y)
	if err != nil {
		t.Errorf("createScatterPlot() returned an error for valid data: %v", err)
	}

	// Test case 2: Empty data
	x = []float64{}
	y = []float64{}
	createScatterPlot(2, x, y)
	if err == nil {
		t.Error("createScatterPlot() did not return an error for empty data")
	}
}

func TestMean(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	meanValue, err := mean(x)
	if err != nil {
		t.Errorf("mean() returned an error for valid data: %v", err)
	}
	if math.Abs(meanValue-3.0) > 1e-6 {
		t.Errorf("mean() returned an incorrect mean value: expected 3.0, got %f", meanValue)
	}

	// Test case 2: Empty data
	x = []float64{}
	meanValue, err = mean(x)
	if err == nil {
		t.Error("mean() did not return an error for empty data")
	}
	if math.Abs(meanValue-0.0) > 1e-6 {
		t.Errorf("mean() did not return 0.0 for empty data: got %f", meanValue)
	}
}

func TestVariance(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	varianceValue, err := variance(x)
	if err != nil {
		t.Errorf("variance() returned an error for valid data: %v", err)
	}
	if math.Abs(varianceValue-2.5) > 1e-6 {
		t.Errorf("variance() returned an incorrect variance value: expected 2.5, got %f", varianceValue)
	}

	// Test case 2: Empty data
	x = []float64{}
	varianceValue, err = variance(x)
	if err == nil {
		t.Error("variance() did not return an error for empty data")
	}
	if math.Abs(varianceValue-0.0) > 1e-6 {
		t.Errorf("variance() did not return 0.0 for empty data: got %f", varianceValue)
	}
}

func TestStdDev(t *testing.T) {
	// Test case 1: Valid data
	x := []float64{1, 2, 3, 4, 5}
	stdDev, err := stdDev(x)
	if err != nil {
		t.Errorf("stdDev() returned an error for valid data: %v", err)
	}
	if math.Abs(stdDev-math.Sqrt(2.5)) > 1e-6 {
		t.Errorf("stdDev() returned an incorrect standard deviation value: expected %f, got %f", math.Sqrt(2.5), stdDev)
	}

}

func TestRound(t *testing.T) {
	// Test case 1: Positive number
	x := 3.14159
	rounded := round(x, 2)
	if math.Abs(rounded-3.14) > 1e-6 {
		t.Errorf("round() returned an incorrect rounded value for positive number: expected 3.14, got %f", rounded)
	}

	// Test case 2: Negative number
	x = -3.14159
	rounded = round(x, 2)
	if math.Abs(rounded+3.14) > 1e-6 {
		t.Errorf("round() returned an incorrect rounded value for negative number: expected -3.14, got %f", rounded)
	}
}
