func balancedStringSplit(s string) int {
    lCount := 0
    rCount := 0
    balancedStrings := 0

    for _,character := range s{
        if character == 'R'{
            rCount++
        } else{
            lCount++
        }
        if lCount == rCount{
            balancedStrings++
            rCount = 0
            lCount = 0
        }
    }

    return balancedStrings
}
