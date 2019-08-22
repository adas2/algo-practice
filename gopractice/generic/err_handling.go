package main

import (
	_ "errors"
	"fmt"
	"math"
)
const (
	lineNo = 10
	colNo = 12
)

type SyntaxError struct {
    Line int
    Col  int
}

func (e *SyntaxError) Error() string {
    return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &SyntaxError{lineNo, colNo}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	/*
	if err != nil {
		//etype := err.(type)
		fmt.Printf("Error %v in line\n", err.Col)
		return
	}
	*/
	switch e := err.(type) {
    	case *SyntaxError:
        	// Do something interesting with e.Line and e.Col.
		fmt.Printf("Error %v in line %d\n", err, e.Line)
		return
 
    	default:
        	fmt.Println(e)
	}
	fmt.Printf("Area of circle %0.2f", area)    
}