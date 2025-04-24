package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "./sample.txt"
	
	numbers, err := readIntegersFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	
	_, inversions := SortAndCountInversions(numbers)
	fmt.Printf("Number of inversions: %d\n", inversions)
}

func readIntegersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert each line to integer
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return numbers, nil
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
