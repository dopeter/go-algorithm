package PkgSort

func ShellSort(array []int){

	length:=len(array)
	shellPointIndex:=1

	//1,4,13,40,121,164
	for{
		if shellPointIndex>=length/3{
			break
		}
		shellPointIndex=3*shellPointIndex+1
	}

	for{
		if shellPointIndex<1{
			break
		}

		for index:=shellPointIndex;index<length;index++{
			for index2:=index; index2>=shellPointIndex && CompareInt(array[index2],array[index2-shellPointIndex])<0; index2-=shellPointIndex{
				ExchangeElem(array,index2,index2-shellPointIndex)
			}
		}

		shellPointIndex=shellPointIndex/3
	}

}