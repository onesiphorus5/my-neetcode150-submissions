func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }

    sHashTb, tHashTb := make(map[uint8]int), make(map[uint8]int)

    for i := 0; i<len(s); i++ {
        sHashTb[s[i]] += 1
        tHashTb[t[i]] += 1
    }
    if len(sHashTb) != len(tHashTb) { return false }

    for k, v := range(sHashTb) {
        if tHashTb[k] != v { return false }
    }

    return true
}