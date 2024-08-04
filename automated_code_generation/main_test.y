%{
package main_test

import (
    "testing"
    "main" // Import the package where main.go is located
    "github.com/montanaflynn/stats"
    "math"
)

// Define a function to call the actual implementations from main
func executeTestCase(identifier string, x, y []float64, expectedResults []float64, t *testing.T) {
    switch identifier {
    case "checkDataQuality":
        if err := main.checkDataQuality(x, y); err != nil {
            t.Errorf("Error in checkDataQuality: %v", err)
        }
    case "linearRegression":
        series, err := main.linearRegression(x, y)
        if err != nil {
            t.Errorf("Error in linearRegression: %v", err)
            return
        }
        intercept := series[0].Y
        slope := series[1].Y
        if !almostEqual(intercept, expectedResults[0]) {
            t.Errorf("Expected intercept %.2f, got %.2f", expectedResults[0], intercept)
        }
        if !almostEqual(slope, expectedResults[1]) {
            t.Errorf("Expected slope %.2f, got %.2f", expectedResults[1], slope)
        }
    case "calculateRSquared":
        intercept := expectedResults[0]
        slope := expectedResults[1]
        rSquared, err := main.calculateRSquared(x, y, intercept, slope)
        if err != nil {
            t.Errorf("Error in calculateRSquared: %v", err)
            return
        }
        if !almostEqual(rSquared, expectedResults[2]) {
            t.Errorf("Expected R-squared %.2f, got %.2f", expectedResults[2], rSquared)
        }
    case "Correlation":
        correlation, err := stats.Correlation(x, y)
        if err != nil {
            t.Errorf("Error in correlation: %v", err)
            return
        }
        if !almostEqual(correlation, expectedResults[0]) {
            t.Errorf("Expected correlation %.2f, got %.2f", expectedResults[0], correlation)
        }
    default:
        t.Errorf("Unknown test case identifier: %s", identifier)
    }
}

func almostEqual(a, b float64) bool {
    const epsilon = 0.0001
    return math.Abs(a-b) < epsilon
}

// Define SymType for token and non-terminal types
type yySymType struct {
    tokenValue string
    numValue   float64
    values     []float64
    string
    float64
    []float64
}

%token <string> IDENTIFIER
%token <float64> NUMBER
%left '+' '-'
%left '*' '/'
%left UMINUS

%%

input:
    /* empty */
    | input test_case
    ;

test_case:
    IDENTIFIER '(' values ')' '\n'  { 
        testCases = append(testCases, func(t *testing.T) { executeTestCase($1, $3[:len($3)/2], $3[len($3)/2:], $3[:len($3)/2], t) })
    }
    | IDENTIFIER '(' ')' '\n'     { 
        testCases = append(testCases, func(t *testing.T) { executeTestCase($1, nil, nil, nil, t) }) 
    }
    ;

values:
    value_list
    | /* empty */  { $$ = nil }
    ;

value_list:
    value_list ',' expression { $$ = append($1, $3) }
    | expression            { $$ = []float64{$1} }
    ;

expression:
    NUMBER               { $$ = $1 }
    | IDENTIFIER         { $$ = $1 }
    | expression '+' expression { $$ = $1 + $3 }
    | expression '-' expression { $$ = $1 - $3 }
    | expression '*' expression { $$ = $1 * $3 }
    | expression '/' expression { $$ = $1 / $3 }
    | '-' expression %prec UMINUS { $$ = -$2 }
    | '(' expression ')' { $$ = $2 }
    ;

%%
