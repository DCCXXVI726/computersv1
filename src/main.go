package main

import (
	"os"
	"fmt"
	"log"
)
//Добавить проверку длинны для args
func main() {
	srt, err := ReduceForm(os.Args[1]);
	if (err != nil) {
		log.Fatal(err)
	}
	fmt.Println(srt)
}
