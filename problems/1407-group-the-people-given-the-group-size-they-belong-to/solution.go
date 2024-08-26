func groupThePeople(groupSizes []int) [][]int {
    temp_group := make(map[int][]int)
    var result [][]int
    
    for i, size := range groupSizes {
        temp_group[size] = append(temp_group[size], i)
        
        if len(temp_group[size]) == size {
            result = append(result, temp_group[size])
            temp_group[size] = []int{}
        }
    }
    return result
}
