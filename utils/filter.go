package utils

func Filter(data []string, fn func(string) bool) []string {
	filtered := make([]string, 0)

	for _, item := range data {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
