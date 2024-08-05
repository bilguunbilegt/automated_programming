package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
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
		}
		// Perform linear regression analysis
		series, err := linearRegression(x, y)
		if err != nil {
			log.Fatalf("Failed to perform linear regression: %v", err)
		}
		m := series[1].Y
		_ = series[0].Y
		c := 0.0 // Replace 0.0 with the appropriate value for 'c'
		r, err := calculateRSquared(x, y, c, m)
		if err != nil {
			log.Fatalf("Failed to calculate R^2: %v", err)
		}
		// Write the analysis results to the file
		fmt.Fprintf(file, "Set %d\n", i)
		fmt.Fprintf(file, "Slope: %.2f\n", m)
		fmt.Fprintf(file, "Intercept: %.2f\n", c)
		fmt.Fprintf(file, "R^2: %.2f\n\n", r)
		// Create a scatter plot for the data
		p := plot.New()
		if err != nil {
			log.Fatalf("Failed to create plot: %v", err)
		}
		if err != nil {
			log.Fatalf("Failed to create plot: %v", err)
		}
		p.Title.Text = fmt.Sprintf("Anscombe's Quartet Set %d", i)
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"
		s, err := plotter.NewScatter(plotter.XYs(xyData(x, y)))
		if err != nil {
			log.Fatalf("Failed to create scatter plot: %v", err)
		}
		p.Add(s)
		// Save the plot to a file
		if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("set%d.png", i)); err != nil {
			log.Fatalf("Failed to save plot: %v", err)
		}
		meanValue, err := mean(x)
		if err != nil {
			log.Fatalf("Failed to calculate mean: %v", err)
		}
		varianceValue, err := variance(x)
		if err != nil {
			log.Fatalf("Failed to calculate variance: %v", err)
		}
		stdDevValue, err := stdDev(x)
		if err != nil {
			log.Fatalf("Failed to calculate standard deviation: %v", err)
		}
		fmt.Fprintf(file, "Mean: %.2f\n", meanValue)
		fmt.Fprintf(file, "Variance: %.2f\n", varianceValue)
		fmt.Fprintf(file, "Standard Deviation: %.2f\n\n", stdDevValue)

	}
}

func checkDataQuality(x, y []float64) error {
	if len(x) != len(y) {
		return fmt.Errorf("mismatched data lengths")
	}
	return nil
}

func xyData(x, y []float64) plotter.XYs {
	pts := make(plotter.XYs, len(x))
	for i := range x {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}

func linearRegression(x, y []float64) (stats.Series, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("mismatched data lengths")
	}
	series := make(stats.Series, 2)
	series, err := linearRegression(x, y)
	if err != nil {
		log.Fatalf("Failed to perform linear regression: %v", err)
	}
	m := series[1].Y
	c := series[0].Y
	if err != nil {
		return nil, err
	}
	series[0] = stats.Coordinate{X: x[0], Y: m*x[0] + m}
	series[1] = stats.Coordinate{X: x[len(x)-1], Y: m*x[len(x)-1] + c}
	return series, nil
}

func calculateRSquared(x, y []float64, intercept, slope float64) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("mismatched data lengths")
	}
	var sumOfSquares, sumOfResiduals float64
	for i := range x {
		yPred := intercept + slope*x[i]
		meanY, err := stats.Mean(y)
		if err != nil {
			log.Fatalf("Failed to calculate mean of y: %v", err)
		}
		sumOfSquares += (y[i] - meanY) * (y[i] - meanY)
		sumOfResiduals += (y[i] - yPred) * (y[i] - yPred)
	}
	rSquared := 1 - sumOfResiduals/sumOfSquares
	return rSquared, nil
}

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
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatalf("Failed to create scatter plot: %v", err)
	}
	s.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(s)
	p.Legend.Add(fmt.Sprintf("Set %d", setNumber), s)
	if err := p.Save(5*vg.Inch, 5*vg.Inch, fmt.Sprintf("anscombe_set_%d.png", setNumber)); err != nil {
		log.Fatalf("Failed to save plot: %v", err)
	}
}

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

func mean(x []float64) (float64, error) {
	if len(x) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x)), nil
}

func variance(x []float64) (float64, error) {
	if len(x) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	m, err := mean(x)
	if err != nil {
		return 0, err
	}
	var sum float64
	for _, v := range x {
		sum += (v - m) * (v - m)
	}
	return sum / float64(len(x)), nil
}

func stdDev(x []float64) (float64, error) {
	v, err := variance(x)
	if err != nil {
		return 0, err
	}
	rounded, err := stats.Round(math.Sqrt(v), 2)
	if err != nil {
		return 0, err
	}
	return rounded, nil
}

func round(x float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(x*shift) / shift
}
