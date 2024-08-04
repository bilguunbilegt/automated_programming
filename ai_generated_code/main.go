package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Define the Anscombe map
var anscombeMap = map[string][]float64{
	"x1": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"y1": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	"x2": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"y2": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	"x3": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"y3": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	"x4": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	"y4": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
}

func main() {
	// Open file for writing
	file, err := os.Create("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for key, xValues := range anscombeMap {
		if len(xValues) == 0 {
			msg := fmt.Sprintf("No data for %s\n", key)
			fmt.Print(msg)
			writer.WriteString(msg)
			continue
		}
		yValues := anscombeMap["y"+key[1:]] // Corresponding y-values

		// Prepare slices for linear regression output
		var slope, intercept float64
		var slopeIntercept []float64 = make([]float64, 2)

		// Perform linear regression
		stat.LinearRegression(xValues, yValues, slopeIntercept, false)
		slope = slopeIntercept[0]
		intercept = slopeIntercept[1]

		msg := fmt.Sprintf("Linear Regression for %s:\n", key)
		msg += fmt.Sprintf("Slope: %.2f\n", slope)
		msg += fmt.Sprintf("Intercept: %.2f\n", intercept)
		fmt.Print(msg)
		writer.WriteString(msg)

		// Calculate R squared
		yPredicted := make([]float64, len(xValues))
		for i, x := range xValues {
			yPredicted[i] = slope*x + intercept
		}

		ssTotal := sumSquares(yValues, mean(yValues))
		ssResidual := sumSquaresResidual(yValues, yPredicted)
		rSquared := 1 - (ssResidual / ssTotal)
		msg = fmt.Sprintf("R-squared: %.2f\n", rSquared)
		fmt.Print(msg)
		writer.WriteString(msg)

		// Data quality check
		if len(xValues) != len(yValues) {
			msg = fmt.Sprintf("Warning: Mismatch in lengths of x and y values for %s\n", key)
			fmt.Print(msg)
			writer.WriteString(msg)
		}

		// Write statistics
		meanX := mean(xValues)
		varianceX := variance(xValues, meanX)
		stdDevX := stdDev(xValues, meanX)
		meanY := mean(yValues)
		varianceY := variance(yValues, meanY)
		stdDevY := stdDev(yValues, meanY)

		msg = fmt.Sprintf("Mean X: %.2f\n", meanX)
		msg += fmt.Sprintf("Variance X: %.2f\n", varianceX)
		msg += fmt.Sprintf("Standard Deviation X: %.2f\n", stdDevX)
		msg += fmt.Sprintf("Mean Y: %.2f\n", meanY)
		msg += fmt.Sprintf("Variance Y: %.2f\n", varianceY)
		msg += fmt.Sprintf("Standard Deviation Y: %.2f\n", stdDevY)
		fmt.Print(msg)
		writer.WriteString(msg)

		// Plotting
		p := plot.New()

		p.Title.Text = "Anscombe's Dataset " + key
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"

		scatterData := make(plotter.XYs, len(xValues))
		for i := range xValues {
			scatterData[i].X = xValues[i]
			scatterData[i].Y = yValues[i]
		}
		scatter, err := plotter.NewScatter(scatterData)
		if err != nil {
			log.Fatal(err)
		}
		p.Add(scatter)

		if err := p.Save(6*vg.Inch, 6*vg.Inch, "scatter_plot_"+key+".png"); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Plot saved as scatter_plot_" + key + ".png")
		writer.WriteString("\n")
	}

	// Flush the buffer
	writer.Flush()
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

func mean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func variance(values []float64, mean float64) float64 {
	var variance float64
	for _, v := range values {
		variance += (v - mean) * (v - mean)
	}
	return variance / float64(len(values)-1)
}

func stdDev(values []float64, mean float64) float64 {
	return math.Sqrt(variance(values, mean))
}
