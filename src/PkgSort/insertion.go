package PkgSort

func InsertionSort(array []int){

	for index,_:=range array{

		for index2:=index; index2>0 && CompareInt(array[index2],array[index2-1])<0; index2--{
			ExchangeElem(array,index2-1,index2)
		}

	}

}