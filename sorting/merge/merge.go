package merge

func Sort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	first := Sort(a[:len(a)/2])
	second := Sort(a[len(a)/2:])
	return merge(first, second)

}
func merge(a []int, b []int) []int {
	final := []int{}
	return final
}
