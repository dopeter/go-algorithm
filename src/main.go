package main

import (
"fmt"
"PkgSort"
)

func main(){
	fmt.Println("main func start")

	array:=[]int {32,44,11}

	PkgSort.InsertionSort(array)

	PkgSort.PrintlnArray(array)
}