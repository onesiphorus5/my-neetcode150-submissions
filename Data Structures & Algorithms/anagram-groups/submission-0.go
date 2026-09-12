import "slices"
func groupAnagrams(strs []string) [][]string {
	hashT := make(map[string][]string);

	for _, s := range(strs) {
		b := []byte(s);
		slices.Sort(b)
		b_s := string(b)

		hashT[b_s] = append(hashT[b_s], s)
	}

	res := [][]string{}
	for _, v := range(hashT) {
		 res = append(res, v)
	}
	return res
}
