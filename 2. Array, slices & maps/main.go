package main

import "fmt"

func main() {
	array()
	slices()
	maps()
}

func array() {
	fmt.Println("--Arrays--")
	// initialize array
	var arr [3]int32
	fmt.Println(arr[0:3])
	arr[0] = 3
	arr[1] = 45
	arr[2] = 213
	fmt.Println(arr[0:3])

	// print addresses
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	// initialize with values
	var newArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(newArr[0:3])

	// shorthand
	intArr := [3]int32{11, 22, 33}
	fmt.Println(intArr[0:3])

	// shorthand and no size mentioned
	newIntArr := [...]int32{10, 20, 30}
	fmt.Println(newIntArr[0:3])

	for idx, value := range intArr {
		fmt.Printf("Index: %v, value %v\n", idx, value)
	}
}

func slices() {
	fmt.Println("--Slices--")
	// no size mentioned, then its a slice
	// initialize slice
	var slice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v & the capacity is %v\n", len(slice), cap(slice))
	fmt.Println(slice[0:3])
	// append might increase the capacity of the slice
	slice = append(slice, 7)
	fmt.Printf("The length is %v & the capacity is %v\n", len(slice), cap(slice))

	// we can append a whole slice to a different one
	var intSlice []int32 = []int32{8, 9}
	intSlice = append(intSlice, slice...)
	fmt.Println(intSlice)

	// make constructor to initialize slice -> type, length, capacity
	var newSlice []int32 = make([]int32, 3, 8)
	fmt.Println(newSlice)
}

func maps() {
	fmt.Println("--Maps--")
	var trailMap map[string]int32 = make(map[string]int32)
	fmt.Println(trailMap)

	ageMap := map[string]int32{"Sudhanshu": 23, "Anushka": 22}
	var nameMap = map[uint8]string{23: "Sudhanshu", 22: "Anushka"}
	fmt.Println(ageMap)
	fmt.Println(nameMap)
	// returns the default value
	fmt.Println(ageMap["Jason"])
	fmt.Println(ageMap["Sudhanshu"])
	delete(nameMap, 23)
	fmt.Println(nameMap)
	// optional value to check if value exists or not
	var age, exist = ageMap["Sudhanshu"]
	if exist {
		fmt.Printf("The age is of Sudhanshu is %v\n", age)
	} else {
		fmt.Printf("The value does not exists\n")
	}

	// iteration
	for key, value := range ageMap {
		fmt.Printf("The %v is %v\n", key, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
