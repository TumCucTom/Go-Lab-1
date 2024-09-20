package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
		for i,item := range slice{
			slice[i] = f(item)
		}
}

// skeleton modified here as arrays are passed by value not reference
func mapArray(f func(a int) int, array [3]int) [3]int {
	for i:=0;i<3;i++{
		array[i] = f(array[i])
	}
	return array
}

func main() {
	var intsArray = [3]int{1,2,3}
	intsArray = mapArray(addOne,intsArray)
	fmt.Println(intsArray)
	intsSlice := []int{1,2,3}
	mapSlice(addOne,intsSlice)
	fmt.Println(intsSlice)
}
