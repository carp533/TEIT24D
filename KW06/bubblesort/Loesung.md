
```go
func BinarySearch(a []int, x int) bool {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := (low + high) / 2
		//zur Veranschaulichung:
        //fmt.Printf("searching from %v to %v, middle element a[%v] = %v\n", low, high, mid, a[mid])
		if a[mid] == x {
			return true
		}
		if a[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
```

```go
//swaped und sorted zum vorzeitigen Verlassen der Schleife kann hier auch eingebaut werden.
func BubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
```
