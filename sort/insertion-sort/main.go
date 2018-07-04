package main

import (
	"fmt"
	"log"
)

func main() {
	array := []int{5, 8 ,9}
	i := search(array, 12)
	if i == -1 {
		log.Fatalln("value not found")
	}
	fmt.Println("index:", i)
}

func insertionSort(arr []int) ([]int) {
	// avoiding the side-effect to make this function
	// pure.
	arrOut := make([]int, len(arr))
	copy(arrOut, arr)
	for j := 1; j < len(arrOut); j++ {
		key := arrOut[j]
		i := j - 1
		for i >= 0 && arrOut[i] > key {
			arrOut[i+1] = arrOut[i]
			i--
		}
		arrOut[i+1] = key
	}
	return arrOut
}

func search(arr []int, val int) (index int) {
	index = -1
	for i, v := range arr {
		if val == v {
			index = i
		}
	}
	return index
}