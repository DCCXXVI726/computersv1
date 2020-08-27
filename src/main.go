package main

import (
	"os"
	"fmt"
	"log"
)
//Добавить проверку длинны для args
func main() {
	str, err := ReduceForm(os.Args[1]);
	if (err != nil) {
		log.Fatal(err)
	}
	fmt.Printf("%f * X^0 + %f * X^1 + %f * X^2 = 0\n", str[0], str[1], str[2])
	FindSolution(str)
}
