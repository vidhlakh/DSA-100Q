func combineMethod2(n, k int) [][]int {
	var res [][]int
	curr := utils.NewStack()
	helperMethod2(1, n, k, &res, curr)
	fmt.Println("res", res)
	return res
}

func helperMethod2(i, n, k int, res *[][]int, curr *utils.Stack) {

	if curr.Len() == k {
		tmp := *curr
		*res = append(*res, tmp)
		fmt.Println("res", *res)
		return
	}
	if i > n {
		return
	}

	// pick
	curr.Push(i)
	helperMethod2(i+1, n, k, res, curr)
	curr.Pop()
	// dont pick
	helperMethod2(i+1, n, k, res, curr)

}