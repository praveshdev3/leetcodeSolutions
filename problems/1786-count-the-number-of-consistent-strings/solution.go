func countConsistentStrings(allowed string, words []string) int {
    allowedChars := make(map[rune]int)
    for _,s := range allowed{
        allowedChars[s]++
    }
    consistentCount := 0
    for _,word := range words{
        isConsistent := true
        for _,w := range word{
            if allowedChars[w] < 1{
                isConsistent = false
                break
            }
        }
        if isConsistent {
            consistentCount++
        }
    }
    return consistentCount
}
