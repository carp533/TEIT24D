package bubble

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleBubbleSort() {
	//Output:
	//[3 8 11 14 17 18 43]
	array := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(BubbleSort(array))
}
func BenchmarkBubbleSort(b *testing.B) {
	size := 100000
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	b.ResetTimer()
	BubbleSort(slice)
}

func ExampleBinarySearch() {
	//Output:
	//Element 17 gefunden.
	a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	x := 17

	if BinarySearch(a, x) {
		fmt.Printf("Element %d gefunden.\n", x)
	} else {
		fmt.Printf("Element %d nicht gefunden.\n", x)
	}

}
