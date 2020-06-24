package practice

// Intersect finds array intersection
func Intersect(nums1 []int, nums2 []int) []int {
	// check for empty cases and return
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	// find shorter array
	var flag bool = false
	var countMap map[int]int
	if len(nums1) < len(nums2) {
		countMap = createMap(nums1)
		flag = true
	} else {
		countMap = createMap(nums2)
	}

	// traverse second array and check intersection
	// create intersection array from map
	var result []int
	if flag {
		result = intersectHelper(nums2, countMap)
	} else {
		result = intersectHelper(nums1, countMap)
	}

	// return intersection array
	return result
}

// create map with array elements as key and counter as value
func createMap(arr []int) map[int]int {
	cMap := map[int]int{}
	for i := range arr {
		if _, ok := cMap[arr[i]]; ok {
			cMap[arr[i]]++
		} else {
			cMap[arr[i]] = 1
		}
	}
	return cMap
}

func intersectHelper(arr []int, cMap map[int]int) []int {
	var out []int
	for j := range arr {
		if val, ok := cMap[arr[j]]; ok && val > 0 {
			out = append(out, arr[j])
			cMap[arr[j]]--
		}
	}
	return out
}
