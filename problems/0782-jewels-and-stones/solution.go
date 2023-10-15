func numJewelsInStones(jewels string, stones string) int {
    jewelsMap := map[rune]bool{}
    for _,jewel := range jewels{
        jewelsMap[jewel] = true
    }

    result := 0

    for _,stone := range stones{
        if jewelsMap[stone] {
            result++
        }
    }

    return result
}
