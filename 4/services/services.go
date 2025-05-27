package services

func ChangeSlices(sl1, sl2 []string) []string {
	res := make([]string, 0)
	map2 := make(map[string]int)
	for _, el := range sl2 {
		map2[el] += 1
	}
	for _, el := range sl1 {
		if _, ok := map2[el]; !ok {
			res = append(res, el)
		}
	}
	return res
}
