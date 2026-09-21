func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
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
			if obstacleGrid[i][j] == 1 { return 0 }
			return 1
		}
		if mem[i][j] != -1 { return mem[i][j] }
		
		// if there is an obstacle at i, j then it's not possible
		// to reach i, j
		if obstacleGrid[i][j] == 1 { return 0 }

		// option1: from the north
		// option2: from the left

		opt1_ret, opt2_ret := 0, 0
		if i > 0 { opt1_ret = helper(i-1, j) }
		if j > 0 { opt2_ret = helper(i, j-1) }

		mem[i][j] = opt1_ret + opt2_ret
		return opt1_ret + opt2_ret
	}

	return helper(m-1, n-1)
}
