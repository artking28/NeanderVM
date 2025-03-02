package main

import (
	"os"
	"log"
	// "fmt"
	"neanderVM/neander"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Missing '.mem' file reference.")
	}

	file := os.Args[1]
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf(err.Error())
	}
	neander.PrintProgram(bytes, false, false)

	// pr, _ := neander.RunProgram(bytes, false, true)
	// fmt.Printf("\nResult:\n\tAc = %x, Pc = %x, Z = %v, N = %v\n", pr.Ac, pr.Pc, pr.Ac == 0, int8(pr.Ac) < 0)
}