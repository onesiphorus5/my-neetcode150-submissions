func uniquePaths(m int, n int) int {
	// ways(i, j) = num of ways to reach i, j from 0, 0
	// base case: i==0, j==0 => 1 way
	// case1: from the north, i-1, j
	// case2: from the est, i, j-1

	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}

	var helper func(i, j int) int
	helper = func(i, j int) int {
		if (i == 0) && (j == 0) { return 1 }

		if mem[i][j] != -1 { return mem[i][j] }

		case1_ret, case2_ret := 0, 0
		if i > 0 { case1_ret = helper(i-1, j) }
		if j > 0 { case2_ret = helper(i, j-1) }

		mem[i][j] = case1_ret + case2_ret
		return mem[i][j]
	}
	return helper(m-1, n-1)
}