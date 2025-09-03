func removeOuterParentheses(s string) string {
    var result string
    var index int

    for _,r := range s{
        switch r{
            case '(' :
                if index > 0{
                    result += string(r)
                }
                index++
            case ')' :
                index--
                if index > 0{
                    result += string(r)
                }
        }
    }

    return result
}
