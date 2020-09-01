package main

import (
	"os"
	"fmt"
	"log"
)
func main() {
	if len(os.Args) == 1 {
		log.Fatal("usage: ./computor \"polynom\"\n")
	}
	i := 1
	for i <  len(os.Args) {
		fmt.Printf("polynom № %d\n", i)
		str, err := ReduceForm(os.Args[i])
		if (err != nil) {
			fmt.Println(err)
		} else {
			k := 2
			flag := 1
			for k >= 0 {
				if str[k] != 0 {
					if flag == 0 {
						fmt.Print(" + ")
					}
					fmt.Printf("%f * X^%d", str[k], k)
					flag = 0
				}
				k--
			}
			if flag == 1 {
				fmt.Println("0 = 0")
			} else {
				fmt.Println(" = 0")
			}
			FindSolution(str)
		}
		i++
	}
}
