func longestConsecutive(nums []int) int {
    mp := make(map[int]int, 0)

    for _,num := range nums{
        mp[num] = 1
    }

    maxSeqLen := 0

    for key := range mp{
        _,ok := mp[key-1]
        if !ok{
            seqLen := 1
            for {
                if _,ok := mp[key+seqLen]; ok{
                    seqLen++
                    continue
                }
                maxSeqLen = max(maxSeqLen, seqLen)
                break
            }
        }
    }

    return maxSeqLen
}
