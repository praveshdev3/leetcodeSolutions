func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i,j int) bool {return points[i][0] < points[j][0]})

    maxWidth := 0
    for i:=1 ; i<len(points); i++{
        maxWidth = max(maxWidth, points[i][0]-points[i-1][0])
    }

    return maxWidth
}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}
