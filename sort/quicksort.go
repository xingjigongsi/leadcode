package sort


//快速排序
func QuickSort(arr []int){

	__quicksort_1(arr,0,len(arr)-1)
}


func __quicksort(arr []int,l,r int){
	if l>=r{
		return
	}
	p:=partion_2(arr,l,r)
	__quicksort(arr,l,p-1)
	__quicksort(arr,p+1,r)
}

//三路快速排序
func __quicksort_1(arr []int,l,r int){
	if l>=r{
		return
	}
	v:= arr[l]
	lt:= l  // [l+1,lt) <v
	gt:=r+1 // [gt,r] >v
	i:= l+1 // [l+1,i)
	for i<gt{
		if v>arr[i]{
			arr[i],arr[lt+1] = arr[lt+1],arr[i]
			lt++
			i++
		}else  if v<arr[i]{
			arr[i],arr[gt-1] = arr[gt-1],arr[i]
			gt--
		}else{
			i++
		}
	}
	// 等于v 的处理
	arr[l],arr[lt] = arr[lt],arr[l]
	lt--
	__quicksort_1(arr,l,lt)
	__quicksort_1(arr,gt,r)
}

// partion 第一部分
func partion(arr []int,l,r int) int {
	v:=arr[l]
	j:=l
	//[l+1,j] [j+1,i)
	for i:=l+1;i<=r;i++{
		if v>arr[i]{
			arr[i],arr[j+1] = arr[j+1],arr[i]
			j++
		}
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}

// [l+1,i) [j+1,r)    partion 第二部分
func partion_2(arr []int,l,r int) int{
	v:=arr[l]
	i,j:=l+1,r
	for{
		for i<=r && v>arr[i]{
			i++
		}
		for j>=l+1 && v<arr[j]{
			j--
		}
		if i>j{
			break
		}
		arr[i],arr[j] = arr[j],arr[i]
		i++
		j--
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}
