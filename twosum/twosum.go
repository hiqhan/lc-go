package main

func twoSum(nums []int, target int) []int {
	checked := make(map[int]int)
	for i, n := range nums {
		_, ok := checked[target-n]
		if ok {
			return []int{checked[target-n], i}
		}
		checked[n] = i
	}
	return nil
}

func main() {
	twoSum([]int{1,3,3}, 6)
}