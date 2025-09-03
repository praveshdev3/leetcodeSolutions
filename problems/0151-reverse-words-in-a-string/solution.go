func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    sArr := strings.Fields(s)
    i := 0
    j := len(sArr) - 1
    for i <= j{
        sArr[i] , sArr[j] = sArr[j], sArr[i]
        i++
        j--
    }
    return strings.Join(sArr," ")
}
