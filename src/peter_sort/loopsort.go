package peter_sort

import(
	"fmt"
)

func LoopSort(array [] int ){
	for index,_:=range array{
		
		//compare
		minIndex:=index
		for index2:=index+1;index2<len(array);index2++ {
			if(CompareInt(array[index2],array[minIndex])<0){
				minIndex=index2
			}
		}

		temp:=array[index]
		array[index]=array[minIndex]
		array[minIndex]=temp
	}

	PrintlnArray(array)
}

func CompareInt(a int,b int) int{
	if a<b{
		return -1
	}else if a==b{
		return 0
	}else{
		return 1
	}
}

func PrintlnArray(array []int){
	for _,val:=range array{
		fmt.Println(val)
	}
	
}