package main

import (
	"fmt"
	"errors"
)

func main(){
	printMe("Jitendra")
	var numerator int = 11
	var denominator int = 2
	var result, rem, err = intDivision(numerator, denominator)
	if err!=nil{
		fmt.Println(err.Error())

	}else if rem == 0{
		fmt.Printf("The result of the integer division is %v\n", result)


	}else{
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, rem)
	}

	// Using switch
	switch{
	case err!=nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, rem)
	}
}

func printMe(name string){
	fmt.Println("Hello! " + name)
}

func intDivision(numerator, denominator int) (int, int, error){
	var err error
	if denominator == 0{
		err = errors.New("Cannot Divide by Zero")
		return 0, 0 ,err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator

	return result, remainder, err
}