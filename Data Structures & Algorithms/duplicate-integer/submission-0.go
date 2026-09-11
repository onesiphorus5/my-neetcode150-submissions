func hasDuplicate(nums []int) bool {
    hashTb := make(map[int]bool)

    for _, n_v := range(nums) {
        if hashTb[n_v] == true { return true }
        hashTb[n_v] = true
    }
    return false
}