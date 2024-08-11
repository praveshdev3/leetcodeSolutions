func findArray(pref []int) []int {
    res := make([]int,len(pref))
    res[0] = pref[0]
    for i:=1; i<len(pref); i++{
        res[i] = pref[i] ^ pref[i-1]
    }
    return res
}
