package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4}
	fmt.Println(arr[3])
	arr[3] = 55
	fmt.Println(arr[3])
	fmt.Println("Array length:", len(arr))

	var slice = []any{1, 2, 3, 4, 5, 6, 7, "\"asdfasdf\"", false}
	fmt.Println(slice)
}
