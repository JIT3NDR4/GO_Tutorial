package main

import "fmt"

func main(){
	// Arrays
	var intArr [3]int32
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// Continuous Memory Allocation
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Default Allocation
	// var intArr2 [3]int32 = [3]int32{1,2,3}
	// intArr2 := [3]int32{1,2,3}
	intArr2  := [...]int32{1,2,3,4}
	fmt.Println((intArr2))

	// Slice
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	
	//Append Another slice using spread
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Create Slice using Make
	var intSlice3 []int32 = make([]int32, 3, 8) // size, capacity
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))


	//MAPS
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	// Loops

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v\n",name, age)
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, value: %v\n", i, v)
	}

	// While Loope
	var i int = 0
	for i <10{
		fmt.Println(i)
		i = i+1
	}

	for i:=0;i<10 ; i++{
		fmt.Println(i)
	}

}