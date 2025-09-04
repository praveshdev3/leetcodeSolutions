func largestOddNumber(num string) string {

    var result string

    for i:=len(num)-1; i>=0; i--{
        n,_ := strconv.Atoi(string(num[i]))
        if n%2 != 0{
            result = num[:i+1]
            break
        }
    }

    return result
}
