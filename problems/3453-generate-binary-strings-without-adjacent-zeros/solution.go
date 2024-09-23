func validStrings(n int) []string {
    var result []string
    result = append(result,getStrings("0", n)...)
    result = append(result,getStrings("1", n)...)
    return result
}

func getStrings(num string, size int) []string{
    var result []string
    if len(num) == size{
        result = append(result,num)
        return result
    }
  
    if num[len(num)-1] == '0'{
        num = num + "1"
        result = append(result,getStrings(num, size)...)
        return result
    } else {
        result = append(result,getStrings(num + "1", size)...)
        result = append(result,getStrings(num + "0", size)...)
        return result
    }
}
