package sort


func selectsort(data []int) []int{
	for i:=0;i<len(data);i++{
		min_index:=i
		for j:= i+1;j<len(data);j++{
			if data[min_index]>data[j]{
				min_index = j
			}
		}
		data[i],data[min_index] = data[min_index],data[i]
	}
	return data
}

//插入排序第一个版本
func insert_1Sort(data []int) []int{
	for i:=1;i<len(data);i++{
		for j:=i;j>0 && data[j]<data[j-1];j--{
			data[j],data[j-1] = data[j-1],data[j]
		}
	}
	return data
}

//插入排序第二个版本
func insert_2sort(data []int) []int{
	for i:=1;i<len(data);i++{
		var j int
		e:=data[i]
		for j =i;j>0 && e<data[j-1];j--{
			data[j],data[j-1] = data[j-1],data[j]
		}
		data[j] = e
	}
	return data
}
