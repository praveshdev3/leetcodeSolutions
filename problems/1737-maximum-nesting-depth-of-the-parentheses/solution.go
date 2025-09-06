func maxDepth(s string) int {
    count := 0
    maxCount := 0
    for i:=0; i<len(s); i++{
        if s[i] == '('{
            count++
        }else if s[i] == ')'{
            if count > maxCount{
                maxCount = count
            }
            count--
        }
    }
    return maxCount
}
