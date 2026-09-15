func rob(nums []int) int {
	n_len := len(nums)

	// rob(i) = max money robbed after house i
	//          case1: rob house i
	//          case2: skip house i
	// 
	// case1: nums[i] + rob(i-2)
	// case2: rob(i-1)
	// base case:
	//      i == 0 => nums[0]
	//      i == 1 => max(nums[0], nums[1])

	mem := make([]int, n_len)
	for i := 0; i < n_len; i++ {
		mem[i] = -1
	}

	var helper func(i int) int
	helper = func(i int) int {
		if i == 0 { return nums[0] }
		if i == 1 { return max(nums[0], nums[1]) }

		if mem[i] != -1 { return mem[i] }

		mem[i] = max(nums[i] + helper(i-2), helper(i-1))
		
		return mem[i]
	}

	return helper(n_len-1)
}
