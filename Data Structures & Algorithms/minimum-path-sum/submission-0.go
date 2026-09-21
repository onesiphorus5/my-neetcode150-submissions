func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)

		for j := 0; j < n; j++ {
			mem[i][j] = -1
		} 
	}

	var helper func(i, j int) int
	helper = func(i, j int) int {
		if (i == 0) && (j == 0) {
			return grid[i][j]
		}
		if mem[i][j] != -1 { return mem[i][j] }
		
		// option1: from the north
		// option2: from the left

		opt1_ret, opt2_ret := -1, -1
		if i > 0 { opt1_ret = helper(i-1, j) + grid[i][j] }
		if j > 0 { opt2_ret = helper(i, j-1) + grid[i][j] }

		if opt1_ret == -1 { 
			mem[i][j] = opt2_ret
		} else if opt2_ret == -1 {
			mem[i][j] = opt1_ret 
		} else {
			mem[i][j] = min(opt1_ret, opt2_ret)
		}
		return mem[i][j]
	}

	return helper(m-1, n-1)
}