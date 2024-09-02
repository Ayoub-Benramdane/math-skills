package main

import (
	"fmt"
	funcs "funcs/Functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Usage : go run . \"PATH TO THE FILE AND FILE NAME\"")
		return
	}
	data := funcs.File(os.Args[1])
	a, b := funcs.LinReg(data)
	c := funcs.CorCoeff(data)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", c)
}
