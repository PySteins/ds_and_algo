package algo

import "github.com/gogf/gf/util/gutil"

func BinarySearch(array []interface{}, v interface{}, comparator gutil.Comparator) int {
	if len(array) == 0 {
		return -1
	}
	begin := 0
	end := len(array)
	for end > begin {
		mid := (end + begin) >> 1
		sub := comparator(array[mid], v)
		switch {
		case sub == 0:
			return mid
		case sub > 0:
			end = mid
		case sub < 0:
			begin = mid + 1
		}
	}
	return -1
}
