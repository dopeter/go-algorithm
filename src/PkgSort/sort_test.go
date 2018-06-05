package PkgSort

import(
	"testing"
	"fmt"
)

func buildTestDescArray() []int{

	length:=10000000
	array:=make([]int,length)

	for index:=0;index<length;index++{
		array[index]=length-index
	}

	return array
}

func TestLoopSort(t *testing.T){
	array:=buildTestDescArray()

	LoopSort(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}

func TestInsertionSort(t *testing.T){
	array:=buildTestDescArray()

	InsertionSort(array)

	//PrintlnArray(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}

func TestShellSort(t *testing.T){

	// this is test that if condition is false , the loop is which continue to increase.
	for loop:=0;loop<100&&loop%2==0;loop++{
		fmt.Println(loop)
	}


	array:=buildTestDescArray()

	ShellSort(array)

	//PrintlnArray(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}