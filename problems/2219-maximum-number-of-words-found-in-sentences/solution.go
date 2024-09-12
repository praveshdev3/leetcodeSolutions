func mostWordsFound(sentences []string) int {
    count := 0
    for _,sentence := range sentences{
        v := strings.Split(sentence, " ")
        if count < len(v){
            count = len(v)
        }
    }
    return count
}
