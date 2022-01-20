package main

import (
	"strings"
	"fmt"
	"strconv"
)

func ReduceForm (str string) (polynom [3]float64, err error) {
	err = nil
	polynom = [3]float64{0, 0, 0}
	terms := strings.Fields(str)
	i := 0
	sign := 1
	equils := 1
	if (terms[0] == "-") {
		sign = -1
		i = 1
	}
	for i < len(terms) {
		number, errLoop := strconv.ParseFloat(terms[i], 64)
		if errLoop != nil {
			err = errLoop
			return polynom, err
		}
		i++
		if (i >= len(terms) || terms[i] != "*") {
			err = fmt.Errorf("every term should respect the form a * X^p")
			return polynom, err
		}
		i++
		if (i >= len(terms) || len(terms[i]) < 3 || terms[i][0:2] != "X^") {
			err = fmt.Errorf("every term should respect the form a * X^p");
			return polynom, err
		}
		degree, errLoop:= strconv.Atoi(terms[i][2:])
		if degree < 0 || degree > 2 {
			err = fmt.Errorf("polynom degree shouldn't be greater than 2.")
			return polynom, err
		}
		polynom[degree] += (float64(sign) * float64(equils) * number)
		if errLoop != nil {
			err = errLoop
			return polynom, err
		}
		i++
		if i >= len(terms) {
			if equils == 1 {
				err = fmt.Errorf("don't have equils mark")
				return polynom, err
			}
			return polynom, nil
		} else if terms[i] == "+" {
			sign = 1
		} else if terms[i] == "-" {
			sign = -1
		} else if terms[i] == "=" && equils != -1 {
			if equils == -1 {
				err = fmt.Errorf("More than 1 equils mark")
				return polynom, err
			}
			equils = -1
			if (i + 1) < len(terms) && terms[i + 1] == "-" {
				sign = -1
				i++
			} else {
				sign = 1
			}
		} else {
			err = fmt.Errorf("every term should respect the form a * X^p")
			return polynom, err
		}
		i++;
	}
	err = fmt.Errorf("unknown err")
	return polynom, err
}
