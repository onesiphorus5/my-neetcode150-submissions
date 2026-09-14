type Solution struct{}

func (s *Solution) Encode(strs []string) string {
// output: for each string first add a byte denoting the number of bytes in the string
    out := []byte{}
    for _, str := range(strs) {
        out = append(out, byte(len(str)))
        out = append(out, []byte(str)...);
    }

    return string(out)
}

func (s *Solution) Decode(encoded string) []string {
    original := []string{}
    
    curr_size := 0
    for i := 0; i<len(encoded); {
        curr_size = int(encoded[i])
        i++ // update at places where a byte is readyy

        str := []byte{}
        end := i + curr_size
        for j := i; j < end; j++ {
            str = append(str, encoded[j]);
            i++ // update at places where a byte is readyy
        }
        original = append(original, string(str));
    }

    return original
}
