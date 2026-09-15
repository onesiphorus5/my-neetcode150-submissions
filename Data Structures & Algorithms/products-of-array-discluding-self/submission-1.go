func productExceptSelf(nums []int) []int {
    res := []int{1}
    res = append(res, nums[:len(nums)-1]...)
    // now res = {1, x1, x2, ... xn-1} 
    p := 1;
    for i, v := range(res) {
        p = p * v;
        res[i] = p
    }
    // now res = {1, x1, x1x2, .. x1x2..xn-1}
    
    p = nums[len(nums)-1]
    for i := len(res)-2; i >=0; i-- {
        res[i] = res[i] * p;
        p = p * nums[i];
    }

    return res
}
