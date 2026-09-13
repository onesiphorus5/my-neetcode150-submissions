type Node struct {
    v, f int
}

type maxheap []Node

func (h maxheap) Less(i, j int) bool { return h[i].f > h[j].f } // for maxheap
func (h maxheap) Len() int {return len(h)}
func (h maxheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h* maxheap) Push(a any) {
    *h = append(*h, a.(Node))
}

func (h* maxheap) Pop() any {
    hv := *h
    r := hv[len(hv)-1];
    *h = hv[:len(hv)-1]
    return r
}

func topKFrequent(nums []int, k int) []int {
    max_heap := maxheap{}

    hashT := make(map[int]int)
    for _, num := range(nums) {
        hashT[num] += 1
    }
    for v, f := range(hashT) {
        max_heap = append(max_heap, Node{v, f})   
    }
    heap.Init(&max_heap)

    res := []int{}
    for i :=0; i<k; i++ {
        // node := heap.Pop(&max_heap).(Node)
        res = append(res, heap.Pop(&max_heap).(Node).v)
    }

    return res
}
