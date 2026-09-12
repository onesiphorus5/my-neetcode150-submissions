func twoSum(nums []int, target int) []int {
    hashT := make(map[int]int);
    for i, v := range(nums) {
        hashT[v] = i
    }
    for j, v := range(nums) {
        c := target - v; // complement
        if i, ok := hashT[c]; ok && i != j {
            return []int{j, i}
        }
    }
    return []int{}
}
