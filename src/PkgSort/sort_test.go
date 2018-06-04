package PkgSort

import(
	"testing"
)

func TestLoopSort(t *testing.T){
	array:=[]int{3,2,1,4}

	LoopSort(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}

func TestInsertionSort(t *testing.T){
	array:=[]int{3,2,1,4,6,6}

	InsertionSort(array)

	PrintlnArray(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}