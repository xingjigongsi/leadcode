package leadcode

import (
	"sort"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	temp := make(map[string][]string)
	result := [][]string{}
	for i := 0; i < len(strs); i++ {
		count := make([]int, 26)
		for _, v := range strs[i] {
			count[v-'a']++
		}
		key := ""
		for _, v := range count {
			key += "#" + strconv.Itoa(v)
		}
		if _, ok := temp[key]; !ok {
			temp[key] = []string{strs[i]}
		} else {
			temp[key] = append(temp[key], strs[i])
		}
	}

	for _, v := range temp {
		result = append(result, v)
	}
	return result
}

// 改进的方法
func groupAnagrams1(strs []string) [][]string {
	temp := make(map[string]int)
	result := [][]string{}
	index := 0
	for i := 0; i < len(strs); i++ {
		count := make([]int, 26)
		for _, v := range strs[i] {
			count[v-'a']++
		}
		key := ""
		for _, v := range count {
			key += "#" + strconv.Itoa(v)
		}
		if idx, ok := temp[key]; !ok {
			temp[key] = index
			result = append(result, []string{strs[i]})
			index++
		} else {
			result[idx] = append(result[idx], strs[i])
		}
	}
	return result
}

func groupAnagrams3(strs []string) [][]string {
	temp := make(map[string]int)
	result := [][]string{}
	index := 0
	for i := 0; i < len(strs); i++ {
		va := strings.Split(strs[i], "")
		sort.Strings(va)
		key := strings.Join(va, "")
		if idx, ok := temp[key]; !ok {
			temp[key] = index
			result = append(result, []string{strs[i]})
			index++
		} else {
			result[idx] = append(result[idx], strs[i])
		}
	}
	return result
}

func main() {

}
