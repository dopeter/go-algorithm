package peter_sort

import(
	"testing"
)

func TestLoopSort(t *testing.T){
	t.Log("start")
	array:=[]int{3,2,1}

	LoopSort(array)

	t.Log("success")
}