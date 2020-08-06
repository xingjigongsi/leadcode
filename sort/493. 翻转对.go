package sort

var result int
func reversePairs(nums []int) int {
	result = 0
	__mergesort(nums,0,len(nums)-1)
	return result
}

func __mergesort(data []int,l,r int){
	if l>=r{
		return
	}
	mid:= (r+l)/2
	__mergesort(data,l,mid)
	__mergesort(data,mid+1,r)
	count:=0
	// 统计
	i:=l
	m:=mid
	j:=m+1
	for i<=m{
		if j>r || data[i]<=2*data[j]{
			i++
			result+=count
		}else{
			j++
			count++
		}
	}
	_merge(data,l,mid,r)

}

func _merge(data []int,l,mid,r int){
	temp:=make([]int,r-l+1)
	for i:=l;i<=r;i++{
		temp[i-l] = data[i]
	}
	i,j:=l,mid+1
	for k:=l;k<=r;k++{
		if i>mid{
			data[k]  = temp[j-l]
			j++
		}else if j>r{
			data[k] = temp[i-l]
			i++
		}else if temp[i-l]<temp[j-l]{
			data[k] = temp[i-l]
			i++
		}else{
			data[k] = temp[j-l]
			j++
		}
	}
}