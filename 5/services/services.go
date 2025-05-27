package services

func InnerSlices(sl1, sl2 []int) (bool, []int) {
	res := make([]int, 0)
	boolean := false
	map2 := make(map[int]int)
	for _, el := range sl2 {
		map2[el] += 1
	}
	for _, el := range sl1 {
		if count, ok := map2[el]; ok && count > 0 {
			map2[el]--
			res = append(res, el)
			boolean = true
		}
	}
	return boolean, res
}
