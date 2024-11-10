package main

import "C"

import "fmt"

//export GoGameDLLInit
func GoGameDLLInit() {
	fmt.Println()
	fmt.Println("GO GameDLLInit")
	fmt.Println()

	fmt.Println("GlobalVars time:", P.GlobalVars.Time())
}
