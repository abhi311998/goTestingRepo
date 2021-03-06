package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Init
	// arr := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Before:", arr)

	// Execution
	BubbleSort(arr)
	fmt.Println("After:", arr)

	// Validation
	if arr[0] != 9 {
		t.Error("Expected", 9, " Got", arr[0], " at index 0")
	}
	if arr[len(arr)-1] != 0 {
		t.Error("Expected", 0, " Got", arr[len(arr)-1], " at last index")
	}
}
