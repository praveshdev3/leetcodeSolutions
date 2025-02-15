func romanToInt(s string) int {
    hm := make(map[uint8]int)
    hm['I'] = 1
    hm['V'] = 5
    hm['X'] = 10
    hm['L'] = 50
    hm['C'] = 100
    hm['D'] = 500
    hm['M'] = 1000

    number := 0
    lastNum := 0
    for i:=len(s)-1; i>=0; i--{
        x,ok := hm[s[i]]; if !ok{
            return number
        }
        if x < lastNum{
            number -= x
        } else {
            number += x
        }
        lastNum = x
    }
    return number
}

