package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	var res [][]int
	N := len(nums)
	helper(nums, 0, &res, N)
	fmt.Printf("Permutation of %v:%v\n", nums, res)
}

func helper(nums []int, ind int, res *[][]int, N int) {

	if ind == N-1 {
		*res = append(*res, nums)
		fmt.Println("after", nums, *res)
		return
	}
	// swap

	for i := ind; i < N; i++ {

		nums[ind], nums[i] = nums[i], nums[ind]
		helper(nums, ind+1, res, N)
		nums[i], nums[ind] = nums[ind], nums[i]

	}

}
