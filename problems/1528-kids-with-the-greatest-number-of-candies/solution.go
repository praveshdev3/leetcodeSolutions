func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool,len(candies))
    maxCandies := 0
    for _,candy := range candies{
        if candy >= maxCandies {
            maxCandies = candy
        }
    }

    for i,candy := range candies{
        if candy + extraCandies >= maxCandies {
            result[i] = true
        }else {
            result[i] = false
        }
    }

    return result
}
