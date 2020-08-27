package main

import (
	"strings"
	"fmt"
	"strconv"
)

func ReduceForm (polynom string) (str string, err error) {
	err = nil
	x := [3]float64{0, 0, 0}
	terms := strings.Fields(polynom)
	i := 0
	sign := 1
	equils := 1
	for i < len(terms) {
		number, errLoop := strconv.ParseFloat(terms[i], 64)
		if errLoop != nil {
			err = errLoop
			return polynom, err
		}
		i++
		if (i >= len(terms) || terms[i] != "*") {
			err = fmt.Errorf("После числа должно идти умножение")
			return polynom, err
		}
		i++
		if (i >= len(terms) || len(terms[i]) < 3 || terms[i][0:2] != "X^") {
			err = fmt.Errorf("После умножения должен идти X в степени");
			return polynom, err
		}
		degree, errLoop:= strconv.Atoi(terms[i][2:])
		if degree < 0 || degree > 3 {
			err = fmt.Errorf("Неправильная степень")
			return polynom, err
		}
		x[degree] += (float64(sign) * float64(equils) * number)
		if errLoop != nil {
			err = errLoop
			return polynom, err
		}
		i++
		if i >= len(terms) {
			fmt.Println(x)
			return polynom, nil
		} else if terms[i] == "+" {
			sign = 1
		} else if terms[i] == "-" {
			sign = -1
		} else if terms[i] == "=" && equils != -1 {
			equils = -1
			sign = 1
		} else {
			err = fmt.Errorf("После степени должно идти + - =")
			return polynom, err
		}
		i++;
		fmt.Println(sign)
	}
	err = fmt.Errorf("Дошел до конца")
	return polynom, err
}
