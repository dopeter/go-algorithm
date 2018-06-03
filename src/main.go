package main

import (
	"fmt"
	"peter_sort"
)

func main(){
	fmt.Println("main func start")

	array:=[]int {32,44,11}
	peter_sort.LoopSort(array)
}