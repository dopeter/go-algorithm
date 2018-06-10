package PkgSort

func Merge(array []int,startIndex int,endIndex int,midIndex int) {
	/**
		两段数组分别是
		start -> mid
		mid+1 -> end
	**/

	leftStartIndex:=startIndex
	rightStartIndex:=midIndex+1

	//copy the array to a temp array
	//todo every time we must copy original array,but original array sometimes maybe has big counts. So we can build copyArray use startIndex and endIndex to build a smaller copyArray
	copyArray:=make([]int,len(array))

	for index:=startIndex;index<=endIndex;index++{
		copyArray[index]=array[index]
	}


	for index:=startIndex;index<=endIndex;index++{
		//left > mid 证明left这一段数字已经排列完成
		if leftStartIndex>midIndex{
			array[index]=copyArray[rightStartIndex]
			rightStartIndex++
			// right > right 证明right这一段数字已经排列完成
		}else  if rightStartIndex>endIndex{
			array[index]=copyArray[leftStartIndex]
			leftStartIndex++
			//比较当前left和right端指向的数字大小
		}else if CompareInt(copyArray[rightStartIndex],copyArray[leftStartIndex])==-1{
			array[index]=copyArray[rightStartIndex]
			rightStartIndex++
			//如果条件都不满足，则按正常遍历顺序排列
		}else{
			array[index]=copyArray[leftStartIndex]
			leftStartIndex++
		}

	}

}


func MergeSortStartFromTop(array []int,startIndex int,endIndex int){
	if endIndex<=startIndex{
		return
	}

	midIndex:=startIndex+(endIndex-startIndex)/2

	MergeSortStartFromTop(array,startIndex,midIndex)
	MergeSortStartFromTop(array,midIndex+1,endIndex)
	Merge(array,startIndex,endIndex,midIndex)
}

func MergeSortStartFromBottom(array[] int){
	arrayLength:=len(array)

	//this func is slow because two reason.
	//1 every time in Merge function will copy orignal array
	//this is loop array
	for childArrayLength:=1;childArrayLength<arrayLength;childArrayLength=childArrayLength+childArrayLength{
		for index:=0;index<arrayLength-childArrayLength;index+=childArrayLength+childArrayLength{
			rightIndex,_:=Min(index+childArrayLength+childArrayLength-1,arrayLength-1)
			midIndex:=index+childArrayLength-1
			Merge(array,index,rightIndex,midIndex)
		}

	}
}
