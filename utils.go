package lab2

// Filter for slice of characters to delete spaces and empty strings
func Filter(arr []string, notSpace func(string) bool) []string {

	var result []string

	// Filtering character slice and appending to result slice
	for i := range arr {
		if notSpace(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result

}