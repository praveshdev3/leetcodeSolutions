func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    runeMap := make(map[rune]int)

    for _,r := range s{
        runeMap[r]++
    }

    for _,r := range t{
        val,ok := runeMap[r]
        if ok && val > 0{
            runeMap[r]--
        }else{
            return false
        }
    }

    return true
}
