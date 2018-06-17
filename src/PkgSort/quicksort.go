package PkgSort

func QuickSort(array []int,start int,end int){
	if(end<=start){
		return
	}

	mid:=partition(array,start,end)

	QuickSort(array,start,mid-1)
	QuickSort(array,mid+1,end)
}

func partition(array []int,start int,end int) int {

	//因为midval是在最后才换到rightDirectionIndex的位置，所以开始比较是从start+1的位置
	leftDirectionIndex:=start+1
	rightDirectionIndex:=end
	midVal:=array[start]

	for{
		for{
			if leftDirectionIndex==end {
				break
			}

			if CompareInt(array[leftDirectionIndex],midVal)<0{
				leftDirectionIndex++
				//一旦遇到比midval大的值，跳出循环，在交换过值后，仍会从当前位置开始往后走
			}else{
				break
			}
		}

		for{
			if rightDirectionIndex==start{
				break
			}

			if CompareInt(array[rightDirectionIndex],midVal)>=0{
				rightDirectionIndex--
			}else{
				break
			}

		}

		if(leftDirectionIndex>=rightDirectionIndex){
			break
		}
		ExchangeElem(array,leftDirectionIndex,rightDirectionIndex)
	}

	ExchangeElem(array,start,rightDirectionIndex)

	return rightDirectionIndex

}

