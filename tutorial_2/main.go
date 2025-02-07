package main

import "fmt"

func main() {
	var intArr = [...]int32{1, 2, 3}
	intArr[1] = 123
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}

	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}

	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Jason is not in the map")
	}
	delete(myMap2, "Adam")

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
