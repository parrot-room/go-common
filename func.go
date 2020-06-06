package common

// OffsetLimit setting default offset and limit for SQL query
// Default limit 10
func OffsetLimit(page, limit int) (int, int) {
	if limit <= 10 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	return offset, limit
}

//UniqueSliceInt32 get unique slice of int32
func UniqueSliceInt32(intSlice []int32) []int32 {
	keys := make(map[int32]bool)
	var list []int32
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UniqueSliceInt get unique slice of int
func UniqueSliceInt(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
