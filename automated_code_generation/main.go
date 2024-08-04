package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var (
	ErrEmptyInput = statsError{"Input must not be empty."}
	ErrNaN        = statsError{"Not a number."}
	ErrNegative   = statsError{"Must not contain negative values."}
	ErrZero       = statsError{"Must not contain zero values."}
	ErrBounds     = statsError{"Input is outside of range."}
	ErrSize       = statsError{"Must be the same length."}
	ErrInfValue   = statsError{"Value is infinite."}
	ErrYCoord     = statsError{"Y Value must be greater than zero."}
)

type statsError struct {
	msg string
}

func (e statsError) Error() string {
	return e.msg
}

func main() {
	// Define the Anscombe dataset
	anscombe := map[string][][]float64{
		"x1": {{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}},
		"y1": {{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}},
		"x2": {{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}},
		"y2": {{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}},
		"x3": {{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}},
		"y3": {{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}},
		"x4": {{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}},
		"y4": {{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}},
	}

	// Creating a file to save the analysis results
	file, err := os.Create("results.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Perform linear regression analysis and print the summary
	for i := 1; i <= 4; i++ {
		x := anscombe[fmt.Sprintf("x%d", i)][0]
		y := anscombe[fmt.Sprintf("y%d", i)][0]

		if err := checkDataQuality(x, y); err != nil {
			log.Printf("Data quality issue in set %d: %v\n", i, err)
			continue
		}

		series, err := linearRegression(x, y)
		if err != nil {
			log.Printf("Error in linear regression for set %d: %v\n", i, err)
			continue
		}

		slope := series[1].Y
		intercept := series[0].Y

		rSquared, err := calculateRSquared(x, y, intercept, slope)
		if err != nil {
			log.Printf("Error in calculating R-squared for set %d: %v\n", i, err)
			continue
		}

		correlation, err := stats.Correlation(x, y)
		if err != nil {
			log.Printf("Error in calculating correlation for set %d: %v\n", i, err)
			continue
		}

		fmt.Printf("Set %d:\n", i)
		fmt.Printf("Intercept: %.2f, Slope: %.2f, R-squared: %.2f, Correlation: %.2f\n\n", intercept, slope, rSquared, correlation)

		result := fmt.Sprintf("Set %d:\nIntercept: %.2f, Slope: %.2f, R-squared: %.2f, Correlation: %.2f\n\n", i, intercept, slope, rSquared, correlation)
		fmt.Printf(result)

		// Write the result to the file
		if _, err := file.WriteString(result); err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
		// Create scatter plot for each dataset
		createScatterPlot(i, x, y)
	}
}

func checkDataQuality(x, y []float64) error {
	if len(x) == 0 || len(y) == 0 {
		return ErrEmptyInput
	}
	if len(x) != len(y) {
		return ErrSize
	}
	for _, val := range x {
		if val != val { // NaN check
			return ErrNaN
		}
	}
	for _, val := range y {
		if val != val { // NaN check
			return ErrNaN
		}
	}
	return nil
}

// func linearRegression(x, y []float64) (intercept, slope, rSquared float64, err error) {
// 	// Perform linear regression
// 	series, err := stats.LinearRegression(makeSeries(x, y))
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	// Extract slope and intercept from the series
// 	slope = series[1].Y
// 	intercept = series[0].Y

// 	// Calculate R-squared
// 	rSquared, err = calculateRSquared(x, y, intercept, slope)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	return intercept, slope, rSquared, nil
// }

func linearRegression(x, y []float64) (stats.Series, error) {
	// Perform linear regression
	series, err := stats.LinearRegression(makeSeries(x, y))
	if err != nil {
		return nil, err
	}
	return series, nil
}

// Helper function to create a series from x and y
func makeSeries(x, y []float64) stats.Series {
	if len(x) != len(y) {
		return nil
	}
	series := make(stats.Series, len(x))
	for i := range x {
		series[i] = stats.Coordinate{X: x[i], Y: y[i]}
	}
	return series
}

// Function to calculate R-squared
func calculateRSquared(x, y []float64, intercept, slope float64) (float64, error) {
	y_mean, err := stats.Mean(y)
	if err != nil {
		return 0, err
	}

	var ss_total, ss_residual float64
	for i := range x {
		y_predicted := intercept + slope*x[i]
		ss_total += (y[i] - y_mean) * (y[i] - y_mean)
		ss_residual += (y[i] - y_predicted) * (y[i] - y_predicted)
	}

	return 1 - (ss_residual / ss_total), nil
}

// Function for scatter plot
func createScatterPlot(setNumber int, x, y []float64) {
	p := plot.New()

	p.Title.Text = fmt.Sprintf("Anscombe's Quartet - Set %d", setNumber)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(x))
	for i := range x {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(scatter)
	p.Legend.Add(fmt.Sprintf("Set %d", setNumber), scatter)

	if err := p.Save(5*vg.Inch, 5*vg.Inch, fmt.Sprintf("anscombe_set_%d.png", setNumber)); err != nil {
		panic(err)
	}
}

//go:generate goyacc -o main_test.go main_test.y
