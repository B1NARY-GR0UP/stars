package hashtable

func isDigit(v int32) bool {
	// 0 - 9
	arr := []int32{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	left, right := 0, 9
	for left <= right {
		mid := (right-left)/2 + left
		num := arr[mid]
		if num == v {
			return true
		} else if num > v {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
