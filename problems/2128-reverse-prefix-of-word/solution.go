func reversePrefix(word string, ch byte) string {
    index := strings.IndexByte(word,ch)
    if index == -1{
        return word
    }

    prefix := word[:index + 1]
    result := reverseString(prefix)
    return result + word[index + 1:]
}

func reverseString(s string) string{
    str := []rune(s)
    for i,j := 0, len(str)-1; i<j; i,j=i+1,j-1{
        str[i], str[j] = str[j], str[i]
    }
    return string(str)
}
