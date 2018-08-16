package PkgSort

import(
	"fmt"
)

//a<b -1; a==b 0;a>b 1
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

func CheckLessSortCompleted(array []int) bool{

	for index,val:=range array{
		if index<len(array)-1 && val>array[index+1]{
			return false
		}
	}

	return true
}

func CheckStackLessSortCompleted(array []int) bool{

	for index,val:=range array{
		if index==0{
			continue
		}

		if index<len(array)-1 && val>array[index+1]{
			return false
		}
	}

	return true
}

func ExchangeElem(array []int ,index1 int,index2 int){
	temp:=array[index1]
	array[index1]=array[index2]
	array[index2]=temp
}

func Min(a int,b int) (int,error){
	if a<b{
		return a,nil
	}else if b<a{
		return b,nil
	}else{
		return a,nil
		//return -1,errors.New("the nums are equal.")
	}
}