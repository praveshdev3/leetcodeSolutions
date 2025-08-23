func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := make([][]int,0)

    for _,interval := range intervals{
        if len(result) == 0 || interval[0] > result[len(result)-1][1]{
            result = append(result, interval)
        } else{
            result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
        }
    }

    return result
}

func max(a,b int)int{
    if a>b {
        return a
    }
    return b
}
