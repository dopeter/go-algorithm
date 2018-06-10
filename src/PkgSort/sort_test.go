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

	// this is test that if condition is false , the loop is which continue to increase.
	for loop:=0;loop<100&&loop%2==0;loop++{
		//fmt.Println(loop)
	}


	array:=buildTestDescArray()

	ShellSort(array)

	//PrintlnArray(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}

}
/*
func TestMerge(t *testing.T){
	array:=[]int {5,6,1,2,3}

	Merge(array,0,4,1)

	PrintlnArray(array)
}
*/
func TestMergeSortStartFromTop(t *testing.T){
	//array:=buildTestDescArray()

	array:=[]int{4,5,6,3,2,1}

	MergeSortStartFromTop(array,0,len(array)-1)

	//PrintlnArray(array)

	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}

func TestMergeSortStartFromBottom(t *testing.T){
	//array:=[]int{4,5,6,3,2,1,7,8}
	array:=buildTestDescArray()
	MergeSortStartFromBottom(array)

	//PrintlnArray(array)
	if CheckLessSortCompleted(array) ==false{
		t.Error("Sort failed.")
	}
}