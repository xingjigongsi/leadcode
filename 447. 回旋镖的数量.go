package leadcode

func numberOfBoomerangs(points [][]int) int {
	result:=0
	for i:=0;i<len(points);i++{
		num_map:=make(map[int]int,0)
		for j:=0;j<len(points);j++{
			if i!=j{
				num_map[dist(points[i],points[j])]++
			}
		}
		for v,_:=range num_map{
			result+=v*(v-1)
		}

	}
	return result
}

//距离
func dist(p1 []int,p2 []int) int{
	return  (p2[1]-p1[1])*(p2[1]-p1[1])+(p2[0]-p1[0])*(p2[0]-p1[0])
}
