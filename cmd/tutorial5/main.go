package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = []rune("résumé")
	var idx = myString[1]
	fmt.Printf("%v, %T\n", idx, idx)

	for i, v:=range myString{
		fmt.Println(i,v)

	}

	var strSlice = []string{"S", "U", "b", "s", "C", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%s\n", catStr)


}