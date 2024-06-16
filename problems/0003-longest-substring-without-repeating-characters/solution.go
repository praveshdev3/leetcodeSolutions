func lengthOfLongestSubstring(s string) int {
    mx := 0
    mp := make(map[byte]int)
    l := 0
    for r:=0; r<len(s); r++{
        mp[s[r]]++
        for mp[s[r]] > 1 {
            mp[s[l]]--
            l++
        }
        mx = max(mx,r-l+1)
    }
    return mx
}
