func stoneGameVI(aliceValues []int, bobValues []int) int {
    indexes := make([]int, len(aliceValues))

    for i := range aliceValues{
        indexes[i] = i
    }

    sort.Slice(indexes, func(i,j int) bool{
        return aliceValues[indexes[i]] + bobValues[indexes[i]] > aliceValues[indexes[j]] + bobValues[indexes[j]]
    })

    res := 0
    for i,v := range indexes{
        if i%2 == 0{
            res+=aliceValues[v]
        }else{
            res-=bobValues[v]
        }
    }

    if res > 0{
        return 1
    } else if res == 0{
        return 0
    }
    return -1
}
