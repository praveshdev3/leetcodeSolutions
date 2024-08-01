import "math"
func findPermutationDifference(s string, t string) int {
    ans := 0
    mp := make(map[rune]int)
    for i,v := range t{
        mp[v] = i
    }

    for i,v := range s{
        ans += int(math.Abs(float64(i-mp[v])))
    }
    return ans
}
