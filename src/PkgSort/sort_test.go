package PkgSort

import(
	"testing"
)

func buildTestDescArray() []int{

	length:=100000
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
	array:=buildTestDescArray()

	ShellSort(array)

	//PrintlnArray(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}