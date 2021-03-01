package sorting

// InsertionSort sorting algorithm implementation
// Time complexity: O(n^2)
// Space complexity: O(1)
func InsertionSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
}
