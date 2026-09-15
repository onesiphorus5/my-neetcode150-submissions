func minCostClimbingStairs(cost []int) int {
	// let's use minCost to shorten minCostClimbingStairs
	// minCost(i): means the minimum cost to move to i+1 floor
	// minCost(i) = min { case1, case2 }
	// case1: 1 step (from i to i+1)    => cost[i] + minCost(i-1)
	// case2: 2 steps (from i-1 to i+1) => cost[i-1] + minCost(i-2)

	c_len := len(cost)

	mem := make([]int, c_len) // TODO: check if len+1
	for i := 0; i<c_len; i++ {
		mem[i] = -1
	}
	var helper func(i int) int
	helper = func(i int) int {
		if i == 0 {
			// this is a special case:
			// to move to 1 we can:
			// 1. start from 1, cost 0
			// 2. start from 0, cost[0]
			return min(0, cost[0])
		}
		if i == 1 { // to move to 2, 2 cases are possible
			return min(cost[0], cost[1])
		}
		if mem[i] != -1 { return mem[i] }

		mem[i] = min(cost[i] + helper(i-1), cost[i-1] + helper(i-2))
		
		return mem[i]
	}

	return helper(c_len-1)
}
