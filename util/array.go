package util

// IsIn 查询ID是否在数组中存在
func IsIn(s []uint, x uint) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}
