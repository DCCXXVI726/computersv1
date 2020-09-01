package main

import (
	"fmt"
	"math"
)

func FindSolution(coef [3]float64) {
	if coef[2] == 0 {
		if coef[1] == 0 {
			if coef[0] == 0 {
				fmt.Println("Solution is field of rational numbers")
			} else {
				fmt.Println("There is no solution")
			}
		} else {
			fmt.Println("Solution is ", coef[0] / coef[1])
		}
	} else {
		disc := coef[1] * coef[1] - 4 * coef[0] * coef[2]
		fmt.Println("Discriminant is ", disc)
		if disc < 0 {
			fmt.Println("There is no solution")
		} else if disc > 0 {
			fmt.Println("Solution are:")
			sqrtDisc := math.Sqrt(disc)
			fmt.Println((-coef[1] + sqrtDisc) / (2 * coef[2]))
			fmt.Println((-coef[1] - sqrtDisc) / (2 * coef[2]))
		} else {
			fmt.Println("Solution is:")
			fmt.Println(-coef[1] / (2 * coef[2]))
		}
	}
}
