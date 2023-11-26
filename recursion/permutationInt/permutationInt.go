package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	var res [][]int
	helper(nums, 0, &res)
	fmt.Printf("Permutation of %v:%v\n", nums, res)
}

func helper(nums []int, ind int, res *[][]int) {
	if ind == len(nums)-1 {
		*res = append(*res, nums)
		fmt.Println(*res)
		return
	}
	// swap

	for i := ind; i < len(nums); i++ {
		nums[ind], nums[i] = nums[i], nums[ind]
		helper(nums, ind+1, res)
		nums[i], nums[ind] = nums[ind], nums[i]

	}

}
