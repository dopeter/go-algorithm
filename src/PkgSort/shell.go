package PkgSort

import(
	"fmt"
)

func ShellSort(array []int){

	length:=len(array)
	shellPointIndex:=1

	//1,4,13,40,121,364
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
		fmt.Println("Current Loop Shell point:",shellPointIndex)
		/**
		Current Loop Shell point: 9841
		Current Loop Shell point: 3280
		Current Loop Shell point: 1093
		Current Loop Shell point: 364
		Current Loop Shell point: 121
		Current Loop Shell point: 40
		Current Loop Shell point: 13
		Current Loop Shell point: 4
		Current Loop Shell point: 1

		插入排序是对整个数组以Index递增为边界且对左侧数组进行排序的方法，因此插入排序对具有一定顺序的数组排序性能较好
		shell排序是预先对数组进行预加工，按照某个步长进行预排序
		例如point为40，则是对数组中所有步长为40的元素进行插入排序
		point为1时，则是对整个数组进行插入排序

		for index2:=index; index2>=shellPointIndex && CompareInt(array[index2],array[index2-shellPointIndex])<0; index2-=shellPointIndex
			index2>=shellPointIndex作为边界点
			CompareInt(array[index2],array[index2-shellPointIndex])<0 判定是否需要调换 当index2==shellPointIndex时，array[index2-shellPointIndex]可以跳到边界点左侧
			index2-=shellPointIndex 每次跳过index%point=0的步数
		**/

		for index:=shellPointIndex;index<length;index++{
			for index2:=index; index2>=shellPointIndex && CompareInt(array[index2],array[index2-shellPointIndex])<0; index2-=shellPointIndex{
				ExchangeElem(array,index2,index2-shellPointIndex)
			}
		}

		shellPointIndex=shellPointIndex/3
	}

}