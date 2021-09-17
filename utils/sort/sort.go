package sort

func BubbleSort(arr []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] < arr[i+1] {
				keepWorking = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
