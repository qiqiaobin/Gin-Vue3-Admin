package slice

// 判断参数一接口是否包含参数二
func Contains(sl []interface{}, v interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// 判断参数一Int是否包含参数二
func ContainsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// 判断参数一Int64是否包含参数二
func ContainsInt64(sl []int64, v int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// 判断参数一字符串是否包含参数二
func ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// 判断参数一切片是否包含参数二
func ContainsSlice(smallSlice, bigSlice []string) bool {
	for i := 0; i < len(smallSlice); i++ {
		if !ContainsString(bigSlice, smallSlice[i]) {
			return false
		}
	}

	return true
}
