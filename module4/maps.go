package main

import (
	"fmt"
	"maps"
)

func main() {

	//creating a map
	m := make(map[string]string)

	//adding element to the map
	m["name"] = "golang"
	m["price"] = "100"
	//getting the element via key
	fmt.Println(m["name"])

	//if the key doesnt exist in map it returns empty vaue
	fmt.Println(m["phone"])

	//printing len of map
	fmt.Println(len(m))

	//deleting a value from the map
	delete(m, "price")
	fmt.Println(m)

	//clearing the map
	// clear(m)

	//creating a map without make function
	mp := map[string]int{"price": 100, "phone": 12345}
	fmt.Println(mp)

	//find function in map
	//val returns the value of the key if found
	//ok returns bool value whether the element is present or not in the map
	val, ok := mp["phone"]
	if ok {
		fmt.Println("Item found with value:=", val)

	} else {
		fmt.Println("Item not found!")
	}

	//checking if two maps are equal with maps inbuilt function
	mp2 := map[string]int{"price": 100, "phone": 12345}

	fmt.Println(maps.Equal(mp, mp2))
}
