package main

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestLinearRegression(t *testing.T) {
	// Sample data
	series := Series{
		{X: 1, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 4},
		{X: 4, Y: 5},
		{X: 5, Y: 6},
	}

	// Perform linear regression
	regressions, err := stats.LinearRegression(series)
	if err != nil {
		t.Fatalf("LinearRegression error: %v", err)
	}

	// Expected results
	expectedSlope := 1.0
	expectedIntercept := 1.0

	if math.Abs(regressions.Slope-expectedSlope) > 1e-6 {
		t.Errorf("Expected slope %f, got %f", expectedSlope, regressions.Slope)
	}
	if math.Abs(regressions.Intercept-expectedIntercept) > 1e-6 {
		t.Errorf("Expected intercept %f, got %f", expectedIntercept, regressions.Intercept)
	}
}

func TestRSquared(t *testing.T) {
	xValues := []float64{1, 2, 3, 4, 5}
	yValues := []float64{2, 3, 4, 5, 6}

	// Perform linear regression
	series := Series{}
	for i := range xValues {
		series = append(series, Coordinate{X: xValues[i], Y: yValues[i]})
	}
	regressions, err := stats.LinearRegression(series)
	if err != nil {
		t.Fatalf("LinearRegression error: %v", err)
	}

	// Predict values
	yPredicted := make([]float64, len(xValues))
	for i, x := range xValues {
		yPredicted[i] = regressions.Slope*x + regressions.Intercept
	}

	// Calculate R-squared
	ssTotal := sumSquares(yValues, mean(yValues))
	ssResidual := sumSquaresResidual(yValues, yPredicted)
	rSquared := 1 - (ssResidual / ssTotal)

	expectedRSquared := 1.0 // Perfect fit
	if math.Abs(rSquared-expectedRSquared) > 1e-6 {
		t.Errorf("Expected R-squared %f, got %f", expectedRSquared, rSquared)
	}
}

func TestDataQuality(t *testing.T) {
	xValues := []float64{1, 2, 3, 4, 5}
	yValues := []float64{2, 3, 4, 5}

	if len(xValues) != len(yValues) {
		t.Error("Mismatch in lengths of x and y values")
	}
}

// Helper functions for testing
func mean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func sumSquares(values []float64, mean float64) float64 {
	var ss float64
	for _, v := range values {
		ss += (v - mean) * (v - mean)
	}
	return ss
}

func sumSquaresResidual(yValues, yPredicted []float64) float64 {
	var ss float64
	for i, y := range yValues {
		ss += (y - yPredicted[i]) * (y - yPredicted[i])
	}
	return ss
}
