package str

import "sort"

// KeysOfMap 字符串Map中key的有序数组
func KeysOfMap(m map[string]string) []string {
	keys := make(sort.StringSlice, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	keys.Sort()
	return []string(keys)
}
