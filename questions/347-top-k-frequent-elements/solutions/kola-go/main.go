package main

import "fmt"


func main() {

}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)
	buckets := make([][]int, len(nums)+1)
	for key, v := range m {
		buckets[v] = append(buckets[v], key)
		fmt.Println(buckets)
	}
	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
		}
	}
	return res
}

