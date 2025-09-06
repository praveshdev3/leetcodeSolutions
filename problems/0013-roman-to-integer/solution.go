func romanToInt(s string) int {
    mp := make(map[byte]int)
    mp['M'] = 1000
    mp['D'] = 500
    mp['C'] = 100
    mp['L'] = 50
    mp['X'] = 10
    mp['V'] = 5
    mp['I'] = 1

    num := 0
    for i:=0; i<len(s); i++{
        if i<len(s)-1{
            val, _ := mp[s[i]]
            val1, _ := mp[s[i+1]]
            if val >= val1{
                num += val
            }else{
                num -= val
            }
        }else{
            val, _ := mp[s[i]]
            num += val
        }
    }

    return num
}
