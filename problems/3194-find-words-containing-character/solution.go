func findWordsContaining(words []string, x byte) []int {
    res := make([]int,0)
    for i,wrd := range words {
        for _,r:=range wrd{
            if byte(r) == x {
                res = append(res,i)
                break
            }
        }
    }
    return res
}
