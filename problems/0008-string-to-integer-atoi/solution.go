func myAtoi(s string) int {
    s = strings.Trim(s, " ")

    start := 0
    signedInteger := 1
    
    if len(s) == 0{
        return 0
    }

    if s[0] == '-'{
        signedInteger = -1
        start = 1
    }else if s[0] == '+'{
        start = 1
    }

    var res int
    for i := start; i < len(s); i++{
        if !(s[i] >= '0' && s[i] <= '9'){
            return res * signedInteger
        }
        res = res*10 + int(s[i]-'0')
        if res*signedInteger <= math.MinInt32{
            return math.MinInt32
        }else if res*signedInteger >= math.MaxInt32{
            return math.MaxInt32
        }
    }
    return res*signedInteger
}
