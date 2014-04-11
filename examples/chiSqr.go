package main

import (
	"fmt"

	"github.com/phil-mansfield/num"
)

func main() {
	nu := 12.0
	chis := []float64{ 6.3038, 8.43842, 11.3403, 14.8454, 18.5493 }

	fmt.Println("Chi squared distribution for nu = %f:", nu)
	fmt.Printf("10th percentile: %g\n", num.IncGamma(nu / 2, chis[0] / 2))
	fmt.Printf("25th percentile: %g\n", num.IncGamma(nu / 2, chis[1] / 2))
	fmt.Printf("50th percentile: %g\n", num.IncGamma(nu / 2, chis[2] / 2))
	fmt.Printf("75th percentile: %g\n", num.IncGamma(nu / 2, chis[3] / 2))
	fmt.Printf("90th percentile: %g\n", num.IncGamma(nu / 2, chis[4] / 2))
}
