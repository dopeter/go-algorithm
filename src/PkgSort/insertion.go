package PkgSort

func InsertionSort(array []int){

	/**
	index递增，以index为边界，对index左侧的数组进行排序
	根据条件CompareInt(array[index2],array[index2-1])<0 判断是否需要调换数字的位置，因为从index=1时就对array[0]和array[1]排好了次序，
	所以当index=2时，如果需要调换就会往左侧走，直到不需调换为止；如果一开始就不满足条件，则直接跳出循环
	**/
	for index,_:=range array{

		for index2:=index; index2>0 && CompareInt(array[index2],array[index2-1])<0; index2--{
			ExchangeElem(array,index2-1,index2)
		}

	}

}