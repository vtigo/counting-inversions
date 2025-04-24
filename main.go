package main
import (
	"fmt"
)
func main() {
	a := []int{6,5,4,3,2,1}
	
	_, inversions := SortAndCountInversions(a)
	fmt.Println(inversions)
}
func SortAndCountInversions(a []int) ([]int, int) {
	n := len(a)
	if n <= 1 {
		return a, 0
	}
	
	b, bCount := SortAndCountInversions(a[:n/2])
	c, cCount := SortAndCountInversions(a[n/2:])
	d, dCount := MergeAndCountSplitInversions(b, c)
	return d, bCount + cCount + dCount
}
func MergeAndCountSplitInversions(b []int, c[]int) ([]int, int) {
	d := make([]int, len(b) + len(c))
	count := 0
	i := 0
	j := 0
	
	for k := range d {
		if i >= len(b) {
			d[k] = c[j]
			j++
		} else if j >= len(c) {
			d[k] = b[i]
			i++
		} else if b[i] <= c[j] {
			d[k] = b[i]
			i++
		} else {
			d[k] = c[j]
			j++
			count += len(b) - i
		}
	}
	return d, count
}
