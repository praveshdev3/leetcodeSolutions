func isPalindrome(s string) bool {
    lowercase := func(r rune)rune{
        if !unicode.IsLetter(r) && !unicode.IsNumber(r){
            return -1
        }
        return unicode.ToLower(r)
    }

    str := strings.Map(lowercase,s)

    for i,j := 0,len(str)-1; i<j; i,j = i+1,j-1{
        if str[i] != str[j]{
            return false
        }
    }

    return true
}
