package PkgSort


func LoopSort(array [] int ){
	for index,_:=range array{
		
		//compare
		minIndex:=index
		for index2:=index+1;index2<len(array);index2++ {
			if(CompareInt(array[index2],array[minIndex])<0){
				minIndex=index2
			}
		}

		ExchangeElem(array,index,minIndex)
	}

	PrintlnArray(array)
}

