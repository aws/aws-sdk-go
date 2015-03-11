package utilsort

import "sort"

func SortedKeys(m map[string]interface{}) []string {
	sorted := []string{}
	for k, _ := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
