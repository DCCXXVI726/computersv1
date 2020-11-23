package main

import (
	"fmt"
)

func Abs(x float64) (float64){
	if (x < 0) {
		return  -1 * x
	}
	return x
}

func Sqrt(x float64) (float64){
	z := 1.0
	delta := 1.0
	acc := 0.00000000001
	for Abs(delta) > acc {
		delta = (z * z - x) / (2 * z)
		z = z - delta
	}
	return z
}

func FindSolution(coef [3]float64) {
	if coef[2] == 0 {
		if coef[1] == 0 {
			if coef[0] == 0 {
				fmt.Println("Solution is field of rational numbers")
			} else {
				fmt.Println("There is no solution")
			}
		} else {
			fmt.Println("Solution is ", -1 * coef[0] / coef[1])
		}
	} else {
		disc := coef[1] * coef[1] - 4 * coef[0] * coef[2]
		fmt.Println("Discriminant is ", disc)
		if disc < 0 {
			fmt.Println("Solution are:")
			re := -1 * coef[1] / (2 * coef[2])
			img := Sqrt(-disc) / (2 * coef[2])
			fmt.Println(re, "+", img, "* i")
			fmt.Println(re, "-", img, "* i")
		} else if disc > 0 {
			fmt.Println("Solution are:")
			sqrtDisc := Sqrt(disc)
			fmt.Println((-coef[1] + sqrtDisc) / (2 * coef[2]))
			fmt.Println((-coef[1] - sqrtDisc) / (2 * coef[2]))
		} else {
			fmt.Println("Solution is:")
			fmt.Println(-coef[1] / (2 * coef[2]))
		}
	}
}
