func numDecodings(s string) int {
	// dec(i) == number of ways to decode s[0..i] inclusive

	// base cases:
	// 1. i == 0; if s[i] == 0 => return 0
	//            else => return 1
	// 2. i < 0; return 1

	// dec(i) has 2 cases (non-base)
	// case1: s[i] is taken alone
	// case2: s[i] is taken with s[i-1]

	// dec(i) = dec(i-1) + dec(i-2); combining both cases

	mem := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		mem[i] = -1
	}

	var helper func(i int) int
	helper = func(i int) int {
		if i < 0 { return 1 }
		if i == 0 {
			if s[i] == '0' { return 0 }
			return 1
		}

		if mem[i] != -1 { return mem[i] }

		case1_ret := 0
		case2_ret := 0

		if s[i] != '0' { case1_ret = helper(i-1) }
		if s[i-1] != '0' {
			if v, err := strconv.Atoi(s[i-1:i+1]); err == nil {
				if v >= 10 && v <= 26 {
					case2_ret = helper(i-2)
				} 
			} 
		}
		mem[i] = case1_ret + case2_ret
		return mem[i]
	}

	return helper(len(s)-1)
}
