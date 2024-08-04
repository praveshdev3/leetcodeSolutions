func maximumWealth(accounts [][]int) int {
    sum := 0
    for _,account := range accounts {
        summ := 0
        for _,money := range account {
            summ += money
        }
        if summ > sum {
            sum = summ
        }
    }
    return sum
}
