package sort

func Mergesort(arr []int) []int{
	 return  merge_helper(arr,0,len(arr)-1)
}

func merge_helper(arr []int,l,r int) []int{
	if l>=r{
		return nil
	}
	mid:= (l+r)/2
	merge_helper(arr,l,mid)
	merge_helper(arr,mid+1,r)
	return merge_sort(arr,l,mid,r)
}

func merge_sort(arr []int,l ,mid,r int) []int{
	aux:=make([]int,r-l+1)
	for i:=l;i<=r;i++{
		aux[i-l] = arr[i]
	}
	i,j := l,mid+1
	for k:=l;k<=r;k++{
		if i>mid{
			arr[k] = aux[j-l]
			j++
		}else if j>r{
			arr[k] = aux[i-l]
			i++
		}else if aux[i-l]<aux[j-l]{
			arr[k] = aux[i-l]
			i++
		}else{
			arr[k] = aux[j-l]
			j++
		}
	}
	return arr
}


// 自底向上
func Mergesort_1(arr []int) []int{
	for size:=1;size<len(arr);size+=size{   //元素个数
		for i:=0;i+size<len(arr);i+=size+size{
			merge_sort(arr,i,i+size-1,sort_min(i+size+size-1,len(arr)-1))
		}
	}
	return arr
}

func sort_min(x,y int) int{
	if x>y{
		return y
	}
	return x
}


