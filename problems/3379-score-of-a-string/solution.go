import "math"
func scoreOfString(s string) int {
    arr := []rune(s)
    var ans float64 = 0
    for i := 0; i<len(arr)-1; i++ {
        ans += math.Abs(float64(arr[i+1] - arr[i]))
    }
    return int(ans)
}
