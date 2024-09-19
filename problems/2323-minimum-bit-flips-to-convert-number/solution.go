func minBitFlips(start int, goal int) int {
    xorResult := start ^ goal
    count := 0
    for xorResult != 0 {
        xorResult &= (xorResult - 1)
        count++
    }
    return count
}
