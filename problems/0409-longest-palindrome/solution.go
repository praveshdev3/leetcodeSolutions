func longestPalindrome(s string) int {
    hs := make(map[rune]int)
    for _,r := range s{
        hs[r] = hs[r] + 1
    }
    length :=0
    for _,val := range hs{
        if val%2 == 0{
            length += val
        } else {
            length += val - 1
        }
    }
    if length < len(s) {
        return length + 1
    }
    return length
}
