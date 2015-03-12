package utilsort

import "sort"

func SortedKeys(m map[string]interface{}) []string {
	i, sorted := 0, make([]string, len(m))
	for k, _ := range m {
		sorted[i] = k
		i++
	}
	sort.Strings(sorted)
	return sorted
}
