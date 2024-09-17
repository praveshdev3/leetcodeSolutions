func getSneakyNumbers(nums []int) []int {
    numsMap := make(map[int]int)
    var res []int

    for _,num := range nums{
        numsMap[num]++
        if numsMap[num] > 1{
            res = append(res,num)
        }
    }

    return res
}
