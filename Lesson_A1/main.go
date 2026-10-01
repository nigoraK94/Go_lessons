package main

import "fmt"

func main() {

	var arr = [] string{"Bob", "Tom", "Hi"}

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
}